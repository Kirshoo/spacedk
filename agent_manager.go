package spacedk

import (
	"sync"
	"fmt"
	"maps"
	"slices"
	"errors"

	"os"
	"bufio"
)

var (
	CurrentAgentDeleteError = errors.New("cannot delete agent, because its currently active")
)

type AgentManager struct {
	mutex sync.RWMutex

	client *Client
	agentToken map[string]string

	activeAgent string
}

func NewAgentManager(c *Client) *AgentManager {
	return &AgentManager{
		client: c,
		agentToken: make(map[string]string),
	}
}

func (m *AgentManager) ActiveAgent() string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.activeAgent
}

func (m *AgentManager) SetActive(agent string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	token, ok := m.agentToken[agent]
	if !ok {
		return AgentNotTrackedError{Agent: agent}
	}

	m.activeAgent = agent
	m.client.setToken(token)

	return nil
}

// Non blocking assignment method. Should be only used internaly
func (m *AgentManager) assignAgent(agent, token string) {
	m.agentToken[agent] = token
}

// Adds a new agent to be tracked by manager.
// If agent already being tracked, its token is overridden with new one
func (m *AgentManager) AddAgent(agent, token string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.assignAgent(agent, token)
}

// Adds a new agent and switches to it
func (m *AgentManager) AddAgentAndSwitch(agent, token string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.assignAgent(agent, token)

	m.activeAgent = agent
	m.client.setToken(token)
}

// Removes agent from being tracked.
// If agent to remove is current agent, will return
// CurrentAgentDeleteError
func (m *AgentManager) DeleteAgent(agent string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if agent == m.activeAgent {
		return CurrentAgentDeleteError
	}

	if _, ok := m.agentToken[agent]; !ok {
		return AgentNotTrackedError{Agent: agent}
	}

	delete(m.agentToken, agent)

	return nil
}

// Returns all agents that are currently tracked
func (m *AgentManager) TrackedAgents() []string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return slices.Collect(maps.Keys(m.agentToken))
}

func (m *AgentManager) SaveAgents(filepath string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("creating file at '%s': %w", filepath, err)
	}
	defer file.Close()

	for agent, token := range m.agentToken {
		_, err := file.WriteString(fmt.Sprintf("%s %s\n", agent, token))
		if err != nil {
			return fmt.Errorf("writing to '%s': %w", filepath, err)
		}
	}

	return nil
}

func (m *AgentManager) LoadAgents(filepath string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("opening file at '%s': %w", filepath, err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var agent string
		var token string

		_, err := fmt.Sscanf(scanner.Text(), "%s %s\n", &agent, &token)
		if err != nil {
			return fmt.Errorf("reading contents of file at '%s': %w", filepath, err)
		}

		m.assignAgent(agent, token)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner has error: %w", err)
	}

	return nil
}
