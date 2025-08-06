package spacedk

type FactionSymbol string
const (
	CosmicFaction FactionSymbol = "COSMIC"
	VoidFaction = "VOID"
	GalacticFaction = "GALACTIC"
	QuantumFaction = "QUANTUM"
	DominionFaction = "DOMINION"
	AstroFaction = "ASTRO"
	CorsairsFaction = "CORSAIRS"
	ObsidianFaction = "OBSIDIAN"
	AegisFaction = "AEGIS"
	UnitedFaction = "UNITED"
	SolitaryFaction = "SOLITARY"
	CobaltFaction = "COBALT"
	OmegaFaction = "OMEGA"
	EchoFaction = "ECHO"
	LordsFaction = "LORDS"
	CultFaction = "CULT"
	AncientsFaction = "ANCIENTS"
	ShadowFaction = "SHADOW"
	EtherealFaction = "ETHEREAL"
)

type FactionTraitSymbol string
const (
	BureaucraticTrait FactionTraitSymbol = "BUREAUCRATIC"
	SecretiveTrait = "SECRETIVE"
	CapitalisticTrait = "CAPITALISTIC"
	IndustriousTrait = "INDUSTRIOUS"
	PeacefulTrait = "PEACEFUL"
	DistrustfulTrait = "DISTRUSTFUL"
	WelcomingTrait = "WELCOMING"
	SmugglersTrait = "SMUGGLERS"
	ScavengersTrait = "SCAVENGERS"
	RebelliousTrait = "REBELLIOUS"
	ExilesTrait = "EXILES"
	PiratesTrait = "PIRATES"
	RaidersTrait = "RAIDERS"
	ClanTrait = "CLAN"
	GuildTrait = "GUILD"
	DominionTrait = "DOMINION"
	FringeTrait = "FRINGE"
	ForsakenTrait = "FORSAKEN"
	IsolatedTrait = "ISOLATED"
	LocalizedTrait = "LOCALIZED"
	EstablishedTrait = "ESTABLISHED"
	NotableTrait = "NOTABLE"
	DominantTrait = "DOMINANT"
	InescapableTrait = "INESCAPABLE"
	InnovativeTrait = "INNOVATIVE"
	BoldTrait = "BOLD"
	VisionaryTrait = "VISIONARY"
	CuriousTrait = "CURIOUS"
	DaringTrait = "DARING"
	ExploratoryTrait = "EXPLORATORY"
	ResourcefulTrait = "RESOURCEFUL"
	FlexibleTrait = "FLEXIBLE"
	CooperativeTrait = "COOPERATIVE"
	UnitedTrait = "UNITED"
	StrategicTrait = "STRATEGIC"
	IntelligentTrait = "INTELLIGENT"
	Research_focusedTrait = "RESEARCH_FOCUSED"
	CollaborativeTrait = "COLLABORATIVE"
	ProgressiveTrait = "PROGRESSIVE"
	MilitaristicTrait = "MILITARISTIC"
	TechnologicallyAdvancedTrait = "TECHNOLOGICALLY_ADVANCED"
	AggressiveTrait = "AGGRESSIVE"
	ImperialisticTrait = "IMPERIALISTIC"
	Treasure_huntersTrait = "TREASURE_HUNTERS"
	DexterousTrait = "DEXTEROUS"
	UnpredictableTrait = "UNPREDICTABLE"
	BrutalTrait = "BRUTAL"
	FleetingTrait = "FLEETING"
	AdaptableTrait = "ADAPTABLE"
	Self_sufficientTrait = "SELF_SUFFICIENT"
	DefensiveTrait = "DEFENSIVE"
	ProudTrait = "PROUD"
	DiverseTrait = "DIVERSE"
	IndependentTrait = "INDEPENDENT"
	SelfInterestedTrait = "SELF_INTERESTED"
	FragmentedTrait = "FRAGMENTED"
	CommercialTrait = "COMMERCIAL"
	Free_marketsTrait = "FREE_MARKETS"
	EntrepreneurialTrait = "ENTREPRENEURIAL"
)

type ContractType string
const (
	ProcurementContract ContractType = "PROCUREMENT"
	TransportContract = "TRANSPORT"
	ShuttleContract = "SHUTTLE"
)

type ItemSymbol string
const (
	PreciousStonesItem ItemSymbol = "PRECIOUS_STONES"
	QuartzSandItem = "QUARTZ_SAND"
	SiliconCrystalsItem = "SILICON_CRYSTALS"
	AmmoniaIceItem = "AMMONIA_ICE"
	LiquidHydrogenItem = "LIQUID_HYDROGEN"
	LiquidNitrogenItem = "LIQUID_NITROGEN"
	IceWaterItem = "ICE_WATER"
	ExoticMatterItem = "EXOTIC_MATTER"
	AdvancedCircuitryItem = "ADVANCED_CIRCUITRY"
	GravitonEmittersItem = "GRAVITON_EMITTERS"
	IronItem = "IRON"
	IronOreItem = "IRON_ORE"
	CopperItem = "COPPER"
	CopperOreItem = "COPPER_ORE"
	AluminumItem = "ALUMINUM"
	AluminumOreItem = "ALUMINUM_ORE"
	SilverItem = "SILVER"
	SilverOreItem = "SILVER_ORE"
	GoldItem = "GOLD"
	GoldOreItem = "GOLD_ORE"
	PlatinumItem = "PLATINUM"
	PlatinumOreItem = "PLATINUM_ORE"
	DiamondsItem = "DIAMONDS"
	UraniteItem = "URANITE"
	UraniteOreItem = "URANITE_ORE"
	MeritiumItem = "MERITIUM"
	MeritiumOreItem = "MERITIUM_ORE"
	HydrocarbonItem = "HYDROCARBON"
	AntimatterItem = "ANTIMATTER"
	FabMatsItem = "FAB_MATS"
	FertilizersItem = "FERTILIZERS"
	FabricsItem = "FABRICS"
	FoodItem = "FOOD"
	JewelryItem = "JEWELRY"
	MachineryItem = "MACHINERY"
	FirearmsItem = "FIREARMS"
	AssaultRiflesItem = "ASSAULT_RIFLES"
	MilitaryEquipmentItem = "MILITARY_EQUIPMENT"
	ExplosivesItem = "EXPLOSIVES"
	LabInstrumentsItem = "LAB_INSTRUMENTS"
	AmmunitionItem = "AMMUNITION"
	ElectronicsItem = "ELECTRONICS"
	ShipPlatingItem = "SHIP_PLATING"
	ShipPartsItem = "SHIP_PARTS"
	EquipmentItem = "EQUIPMENT"
	FuelItem = "FUEL"
	MedicineItem = "MEDICINE"
	DrugsItem = "DRUGS"
	ClothingItem = "CLOTHING"
	MicroprocessorsItem = "MICROPROCESSORS"
	PlasticsItem = "PLASTICS"
	PolynucleotidesItem = "POLYNUCLEOTIDES"
	BiocompositesItem = "BIOCOMPOSITES"
	QuantumStabilizersItem = "QUANTUM_STABILIZERS"
	NanobotsItem = "NANOBOTS"
	AiMainframesItem = "AI_MAINFRAMES"
	QuantumDrivesItem = "QUANTUM_DRIVES"
	RoboticDronesItem = "ROBOTIC_DRONES"
	CyberImplantsItem = "CYBER_IMPLANTS"
	GeneTherapeuticsItem = "GENE_THERAPEUTICS"
	NeuralChipsItem = "NEURAL_CHIPS"
	MoodRegulatorsItem = "MOOD_REGULATORS"
	ViralAgentsItem = "VIRAL_AGENTS"
	MicroFusion_generatorsItem = "MICRO_FUSION_GENERATORS"
	SupergrainsItem = "SUPERGRAINS"
	LaserRiflesItem = "LASER_RIFLES"
	HolographicsItem = "HOLOGRAPHICS"
	ShipSalvageItem = "SHIP_SALVAGE"
	RelicTechItem = "RELIC_TECH"
	NovelLifeformsItem = "NOVEL_LIFEFORMS"
	BotanicalSpecimensItem = "BOTANICAL_SPECIMENS"
	CulturalArtifactsItem = "CULTURAL_ARTIFACTS"

	// Frames
	FrameProbeItem = "FRAME_PROBE"
	FrameDroneItem = "FRAME_DRONE"
	FrameInterceptorItem = "FRAME_INTERCEPTOR"
	FrameRacerItem = "FRAME_RACER"
	FrameFighterItem = "FRAME_FIGHTER"
	FrameFrigateItem = "FRAME_FRIGATE"
	FrameShuttleItem = "FRAME_SHUTTLE"
	FrameExplorerItem = "FRAME_EXPLORER"
	FrameMinerItem = "FRAME_MINER"
	FrameLightFreighterItem = "FRAME_LIGHT_FREIGHTER"
	FrameHeavyFreighterItem = "FRAME_HEAVY_FREIGHTER"
	FrameTransportItem = "FRAME_TRANSPORT"
	FrameDestroyerItem = "FRAME_DESTROYER"
	FrameCruiserItem = "FRAME_CRUISER"
	FrameCarrierItem = "FRAME_CARRIER"
	FrameBulkFreighterItem = "FRAME_BULK_FREIGHTER"

	// Reactors
	ReactorSolar1Item = "REACTOR_SOLAR_I"
	ReactorFusion1Item = "REACTOR_FUSION_I"
	ReactorFission1Item = "REACTOR_FISSION_I"
	ReactorChemical1Item = "REACTOR_CHEMICAL_I"
	ReactorAntimatter1Item = "REACTOR_ANTIMATTER_I"

	// Engines
	EngineImpulseDrive1Item = "ENGINE_IMPULSE_DRIVE_I"
	EngineIonDrive1Item = "ENGINE_ION_DRIVE_I"
	EngineIonDrive2Item = "ENGINE_ION_DRIVE_II"
	EngineHyperDrive1Item = "ENGINE_HYPER_DRIVE_I"

	// Modules
	ModuleMineralProcessor1Item = "MODULE_MINERAL_PROCESSOR_I"
	ModuleGasProcessor1Item = "MODULE_GAS_PROCESSOR_1"
	ModuleCargoHold1Item = "MODULE_CARGO_HOLD_I"
	ModuleCargoHold2Item = "MODULE_CARGO_HOLD_II"
	ModuleCargoHold3Item = "MODULE_CARGO_HOLD_III"
	ModuleCrewQuarters1Item = "MODULE_CREW_QUARTERS_I"
	ModuleEnvoyQuarters1Item = "MODULE_ENVOY_QUARTERS_I"
	ModulePassengerCabin1Item = "MODULE_PASSENGER_CABIN_I"
	ModuleMicroRefinery1Item = "MODULE_MICRO_REFINERY_I"
	ModuleScienceLab1Item = "MODULE_SCIENCE_LAB_I"
	ModuleJumpDrive1Item = "MODULE_JUMP_DRIVE_I"
	ModuleJumpDrive2Item = "MODULE_JUMP_DRIVE_II"
	ModuleJumpDrive3Item = "MODULE_JUMP_DRIVE_III"
	ModuleWarpDrive1Item = "MODULE_WARP_DRIVE_I"
	ModuleWarpDrive2Item = "MODULE_WARP_DRIVE_II"
	ModuleWarpDrive3Item = "MODULE_WARP_DRIVE_III"
	ModuleShieldGenerator1Item = "MODULE_SHIELD_GENERATOR_I"
	ModuleShieldGenerator2Item = "MODULE_SHIELD_GENERATOR_II"
	ModuleOreRefinery1Item = "MODULE_ORE_REFINERY_I"
	ModuleFuelRefinery2Item = "MODULE_FUEL_REFINERY_II"

	// Mounts
	MountGasSiphon1Item = "MOUNT_GAS_SIPHON_I"
	MountGasSiphon2Item = "MOUNT_GAS_SIPHON_II"
	MountGasSiphon3Item = "MOUNT_GAS_SIPHON_III"
	MountSurveyor1Item = "MOUNT_SURVEYOR_I"
	MountSurveyor2Item = "MOUNT_SURVEYOR_II"
	MountSurveyor3Item = "MOUNT_SURVEYOR_III"
	MountSensorArray1Item = "MOUNT_SENSOR_ARRAY_I"
	MountSensorArray2Item = "MOUNT_SENSOR_ARRAY_II"
	MountSensorArray3Item = "MOUNT_SENSOR_ARRAY_III"
	MountMiningLaser1Item = "MOUNT_MINING_LASER_I"
	MountMiningLaser2Item = "MOUNT_MINING_LASER_II"
	MountMiningLaser3Item = "MOUNT_MINING_LASER_III"
	MountLaserCannon1Item = "MOUNT_LASER_CANNON_I"
	MountMissileLauncher1Item = "MOUNT_MISSILE_LAUNCHER_I"
	MountTurret1Item = "MOUNT_TURRET_I"

	// ShipTypes
	ShipProbeItem = "SHIP_PROBE"
	ShipMiningDroneItem = "SHIP_MINING_DRONE"
	ShipSiphonDroneItem = "SHIP_SIPHON_DRONE"
	ShipInterceptorItem = "SHIP_INTERCEPTOR"
	ShipLightHaulerItem = "SHIP_LIGHT_HAULER"
	ShipCommandFrigateItem = "SHIP_COMMAND_FRIGATE"
	ShipExplorerItem = "SHIP_EXPLORER"
	ShipHeavyFreighterItem = "SHIP_HEAVY_FREIGHTER"
	ShipLightShuttleItem = "SHIP_LIGHT_SHUTTLE"
	ShipOreHoundItem = "SHIP_ORE_HOUND"
	ShipRefiningFreighterItem = "SHIP_REFINING_FREIGHTER"
	ShipSurveyorItem = "SHIP_SURVEYOR"
	ShipBulkFreighterItem = "SHIP_BULK_FREIGHTER"
)

type RawMaterial string
const (	
	PreciousStonesMaterial RawMaterial = RawMaterial(PreciousStonesItem)
	QuartzSandMaterial = RawMaterial(QuartzSandItem)
	SiliconCrystalsMaterial = RawMaterial(SiliconCrystalsItem)
	AmmoniaIceMaterial = RawMaterial(AmmoniaIceItem)
	IceWaterMaterial = RawMaterial(IceWaterItem)
	IronOreMaterial = RawMaterial(IronOreItem)
	CopperOreMaterial = RawMaterial(CopperOreItem)
	AluminumOreMaterial = RawMaterial(AluminumOreItem)
	SilverOreMaterial = RawMaterial(SilverOreItem)
	GoldOreMaterial = RawMaterial(GoldOreItem)
	PlatinumOreMaterial = RawMaterial(PlatinumOreItem)
	DiamondsMaterial = RawMaterial(DiamondsItem)
	UraniteOreMaterial = RawMaterial(UraniteOreItem)
	MeritiumOreMaterial = RawMaterial(MeritiumOreItem)
)

type ProductItem string
const (
	IronProductItem ProductItem = ProductItem(IronItem)
	CopperProductItem = ProductItem(CopperItem)
	SilverProductItem = ProductItem(SilverItem)
	GoldProductItem = ProductItem(GoldItem)
	AluminumProductItem = ProductItem(AluminumItem)
	PlatinumProductItem = ProductItem(PlatinumItem)
	UraniteProductItem = ProductItem(UraniteItem)
	MeritiumProductItem = ProductItem(MeritiumItem)
	FuelProductItem = ProductItem(FuelItem)
)

type FrameType string
const (
	ProbeFrameType FrameType = FrameType(FrameProbeItem)
	DroneFrameType = FrameType(FrameDroneItem)
	InterceptorFrameType = FrameType(FrameInterceptorItem)
	RacerFrameType = FrameType(FrameRacerItem)
	FighterFrameType = FrameType(FrameFighterItem)
	FrigateFrameType = FrameType(FrameFrigateItem)
	ShuttleFrameType = FrameType(FrameShuttleItem)
	ExplorerFrameType = FrameType(FrameExplorerItem)
	MinerFrameType = FrameType(FrameMinerItem)
	LightFreighterFrameType = FrameType(FrameLightFreighterItem)
	HeavyFreighterFrameType = FrameType(FrameHeavyFreighterItem)
	TransportFrameType = FrameType(FrameTransportItem)
	DestroyerFrameType = FrameType(FrameDestroyerItem)
	CruiserFrameType = FrameType(FrameCruiserItem)
	CarrierFrameType = FrameType(FrameCarrierItem)
	BulkFreighterFrameType = FrameType(FrameBulkFreighterItem)
)

type ReactorType string
const (
	Mk1SolarReactorType ReactorType = ReactorType(ReactorSolar1Item)
	Mk1FusionReactorType = ReactorType(ReactorFusion1Item)
	Mk1FissionReactorType = ReactorType(ReactorFission1Item)
	Mk1ChemicalReactorType = ReactorType(ReactorChemical1Item)
	Mk1AntimatterReactorType = ReactorType(ReactorAntimatter1Item)
)

type EngineType string
const (
	Mk1ImpulseDriveEngineType EngineType = EngineType(EngineImpulseDrive1Item)
	Mk1IonDriveEngineType = EngineType(EngineIonDrive1Item)
	Mk2IonDriveEngineType = EngineType(EngineIonDrive2Item)
	Mk1HyperDriveEngineType = EngineType(EngineHyperDrive1Item)
)

type ModuleType string
const (
	Mk1MineralProcessorModuleType ModuleType = ModuleType(ModuleMineralProcessor1Item)
	Mk1GasProcessorModuleType = ModuleType(ModuleGasProcessor1Item)
	Mk1CargoHoldModuleType = ModuleType(ModuleCargoHold1Item)
	Mk2CargoHoldModuleType = ModuleType(ModuleCargoHold2Item)
	Mk3CargoHoldModuleType = ModuleType(ModuleCargoHold3Item)
	Mk1CrewQuartersModuleType = ModuleType(ModuleCrewQuarters1Item)
	Mk1EnvoyQuartersModuleType = ModuleType(ModuleEnvoyQuarters1Item)
	Mk1PassengerCabinModuleType = ModuleType(ModulePassengerCabin1Item)
	Mk1MicroRefineryModuleType = ModuleType(ModuleMicroRefinery1Item)
	Mk1ScienceLabModuleType = ModuleType(ModuleScienceLab1Item)
	Mk1JumpDriveModuleType = ModuleType(ModuleJumpDrive1Item)
	Mk2JumpDriveModuleType = ModuleType(ModuleJumpDrive2Item)
	Mk3JumpDriveModuleType = ModuleType(ModuleJumpDrive3Item)
	Mk1WarpDriveModuleType = ModuleType(ModuleWarpDrive1Item)
	Mk2WarpDriveModuleType = ModuleType(ModuleWarpDrive2Item)
	Mk3WarpDriveModuleType = ModuleType(ModuleWarpDrive3Item)
	Mk1ShieldGeneratorModuleType = ModuleType(ModuleShieldGenerator1Item)
	Mk2ShieldGeneratorModuleType = ModuleType(ModuleShieldGenerator2Item)
	Mk1OreRefineryModuleType = ModuleType(ModuleOreRefinery1Item)
	Mk2FuelRefineryModuleType = ModuleType(ModuleFuelRefinery2Item)
)

type MountType string
const (
	Mk1GasSiphonMountType MountType = MountType(MountGasSiphon1Item)
	Mk2GasSiphonMountType = MountType(MountGasSiphon2Item)
	Mk3GasSiphonMountType = MountType(MountGasSiphon3Item)
	Mk1SurveyorMountType = MountType(MountSurveyor1Item)
	Mk2SurveyorMountType = MountType(MountSurveyor2Item)
	Mk3SurveyorMountType = MountType(MountSurveyor3Item)
	Mk1SensorArrayMountType = MountType(MountSensorArray1Item)
	Mk2SensorArrayMountType = MountType(MountSensorArray2Item)
	Mk3SensorArrayMountType = MountType(MountSensorArray3Item)
	Mk1MiningLaserMountType = MountType(MountMiningLaser1Item)
	Mk2MiningLaserMountType = MountType(MountMiningLaser2Item)
	Mk3MiningLaserMountType = MountType(MountMiningLaser3Item)
	Mk1LaserCannonMountType = MountType(MountLaserCannon1Item)
	Mk1MissileLauncherMountType = MountType(MountMissileLauncher1Item)
	Mk1TurretMountType = MountType(MountTurret1Item)
)

type ShipType string
const (
	ProbeShipType ShipType = ShipType(ShipProbeItem)
	MiningDroneShipType = ShipType(ShipMiningDroneItem)
	SiphonDroneShipType = ShipType(ShipSiphonDroneItem)
	InterceptorShipType = ShipType(ShipInterceptorItem)
	LightHaulerShipType = ShipType(ShipLightHaulerItem)
	CommandFrigateShipType = ShipType(ShipCommandFrigateItem)
	ExplorerShipType = ShipType(ShipExplorerItem)
	HeavyFreighterShipType = ShipType(ShipHeavyFreighterItem)
	LightShuttleShipType = ShipType(ShipLightShuttleItem)
	OreHoundShipType = ShipType(ShipOreHoundItem)
	RefiningFreighterShipType = ShipType(ShipRefiningFreighterItem)
	SurveyorShipType = ShipType(ShipSurveyorItem)
	BulkFreighterShipType = ShipType(ShipBulkFreighterItem)
)

type ShipRole string
const (
	FabricatorRole ShipRole = "FABRICATOR"
	HarvesterRole = "HARVESTER"
	HaulerRole = "HAULER"
	InterceptorRole = "INTERCEPTOR"
	ExcavatorRole = "EXCAVATOR"
	TransportRole = "TRANSPORT"
	RepairRole = "REPAIR"
	SurveyorRole = "SURVEYOR"
	CommandRole = "COMMAND"
	CarrierRole = "CARRIER"
	PatrolRole = "PATROL"
	SatelliteRole = "SATELLITE"
	ExplorerRole = "EXPLORER"
	RefineryRole = "REFINERY"
)

type NavStatus string
const (
	InTransitNavStatus NavStatus = "IN_TRANSIT"
	InOrbitNavStatus = "IN_ORBIT"
	DockedNavStatus = "DOCKED"
)

type FlightMode string
const (
	DriftFlightMode FlightMode = "DRIFT"
	StealthFlightMode = "STEALTH"
	CruiseFlightMode = "CRUISE"
	BurnFlightMode = "BURN"
)

type CrewRotation string
const (
	StrictRotation CrewRotation = "STRICT"
	RelaxedRotation = "RELAXED"
)

type ConditionEventSymbol string
const (
	ReactorOverloadConditionEvent ConditionEventSymbol = "REACTOR_OVERLOAD"
	EnergySpikeFromMineralConditionEvent = "ENERGY_SPIKE_FROM_MINERAL"
	SolarFlareInterferenceConditionEvent = "SOLAR_FLARE_INTERFERENCE"
	CoolantLeakConditionEvent = "COOLANT_LEAK"
	PowerDistributionFluctuationConditionEvent = "POWER_DISTRIBUTION_FLUCTUATION"
	MagneticFieldDisruptionConditionEvent = "MAGNETIC_FIELD_DISRUPTION"
	HullMicrometeoriteStrikesConditionEvent = "HULL_MICROMETEORITE_STRIKES"
	StructuralStressFracturesConditionEvent = "STRUCTURAL_STRESS_FRACTURES"
	CorrosiveMineralContaminationConditionEvent = "CORROSIVE_MINERAL_CONTAMINATION"
	ThermalExpansionMismatchConditionEvent = "THERMAL_EXPANSION_MISMATCH"
	VibrationDamageFromDrillingConditionEvent = "VIBRATION_DAMAGE_FROM_DRILLING"
	ElectromagneticFieldInterferenceConditionEvent = "ELECTROMAGNETIC_FIELD_INTERFERENCE"
	ImpactWithExtractedDebrisConditionEvent = "IMPACT_WITH_EXTRACTED_DEBRIS"
	FuelEfficiencyDegradationConditionEvent = "FUEL_EFFICIENCY_DEGRADATION"
	CoolantSystemAgeingConditionEvent = "COOLANT_SYSTEM_AGEING"
	DustMicroabrasionsConditionEvent = "DUST_MICROABRASIONS"
	ThrusterNozzleWearConditionEvent = "THRUSTER_NOZZLE_WEAR"
	ExhaustPortCloggingConditionEvent = "EXHAUST_PORT_CLOGGING"
	BearingLubricationFadeConditionEvent = "BEARING_LUBRICATION_FADE"
	SensorCalibrationDriftConditionEvent = "SENSOR_CALIBRATION_DRIFT"
	HullMicrometeoriteDamageConditionEvent = "HULL_MICROMETEORITE_DAMAGE"
	SpaceDebrisCollisionConditionEvent = "SPACE_DEBRIS_COLLISION"
	ThermalStressConditionEvent = "THERMAL_STRESS"
	VibrationOverloadConditionEvent = "VIBRATION_OVERLOAD"
	PressureDifferentialStressConditionEvent = "PRESSURE_DIFFERENTIAL_STRESS"
	ElectromagneticSurgeEffectsConditionEvent = "ELECTROMAGNETIC_SURGE_EFFECTS"
	AtmosphericEntryHeatConditionEvent = "ATMOSPHERIC_ENTRY_HEAT"
)

type ComponentSymbol string
const (
	FrameComponent ComponentSymbol = "FRAME"
	ReactorComponent = "REACTOR"
	EngineComponent = "ENGINE"
)

type SystemType string
const (
	NeutronStarSystemType = "NEUTRON_STAR"
	RedStarSystemType = "RED_STAR"
	OrangeStarSystemType = "ORANGE_STAR"
	BlueStarSystemType = "BLUE_STAR"
	YoungStarSystemType = "YOUNG_STAR"
	WhiteDwarfSystemType = "WHITE_DWARF"
	BlackHoleSystemType = "BLACK_HOLE"
	HypergiantSystemType = "HYPERGIANT"
	NebulaSystemType = "NEBULA"
	UnstableSystemType = "UNSTABLE"
)

type WaypointType string
const (
	PlanetWaypointType WaypointType = "PLANET"
	GasGiantWaypointType = "GAS_GIANT"
	MoonWaypointType = "MOON"
	OrbitalStationWaypointType = "ORBITAL_STATION"
	Jump_gateWaypointType = "JUMP_GATE"
	AsteroidFieldWaypointType = "ASTEROID_FIELD"
	AsteroidWaypointType = "ASTEROID"
	EngineeredAsteroidWaypointType = "ENGINEERED_ASTEROID"
	AsteroidBaseWaypointType = "ASTEROID_BASE"
	NebulaWaypointType = "NEBULA"
	DebrisFieldWaypointType = "DEBRIS_FIELD"
	GravityWellWaypointType = "GRAVITY_WELL"
	ArtificialGravityWellWaypointType = "ARTIFICIAL_GRAVITY_WELL"
	FuelStationWaypointType = "FUEL_STATION"
)

type WaypointTraitSymbol string
const (
	UnchartedWaypointTrait WaypointTraitSymbol = "UNCHARTED"
	UnderConstructionWaypointTrait = "UNDER_CONSTRUCTION"
	MarketplaceWaypointTrait = "MARKETPLACE"
	ShipyardWaypointTrait = "SHIPYARD"
	OutpostWaypointTrait = "OUTPOST"
	ScatteredSettlementsWaypointTrait = "SCATTERED_SETTLEMENTS"
	SprawlingCitiesWaypointTrait = "SPRAWLING_CITIES"
	MegaStructuresWaypointTrait = "MEGA_STRUCTURES"
	PirateBaseWaypointTrait = "PIRATE_BASE"
	OvercrowdedWaypointTrait = "OVERCROWDED"
	HighTechWaypointTrait = "HIGH_TECH"
	CorruptWaypointTrait = "CORRUPT"
	BureaucraticWaypointTrait = WaypointTraitSymbol(BureaucraticTrait)
	TradingHubWaypointTrait = "TRADING_HUB"
	IndustrialWaypointTrait = "INDUSTRIAL"
	BlackMarketWaypointTrait = "BLACK_MARKET"
	ResearchFacilityWaypointTrait = "RESEARCHFACILITY"
	MilitaryBaseWaypointTrait = "MILITARY_BASE"
	SurveillanceOutpostWaypointTrait = "SURVEILLANCE_OUTPOST"
	ExplorationOutpostWaypointTrait = "EXPLORATION_OUTPOST"
	MineralDepositsWaypointTrait = "MINERAL_DEPOSITS"
	CommonMetalDepositsWaypointTrait = "COMMON_METAL_DEPOSITS"
	PreciousMetalDepositsWaypointTrait = "PRECIOUS_METAL_DEPOSITS"
	RareMetalDepositsWaypointTrait = "RARE_METAL_DEPOSITS"
	MethanePoolsWaypointTrait = "METHANE_POOLS"
	IceCrystalsWaypointTrait = "ICE_CRYSTALS"
	ExplosiveGasesWaypointTrait = "EXPLOSIVE_GASES"
	StrongMagnetosphereWaypointTrait = "STRONG_MAGNETOSPHERE"
	VibrantAurorasWaypointTrait = "VIBRANT_AURORAS"
	SaltFlatsWaypointTrait = "SALT_FLATS"
	CanyonsWaypointTrait = "CANYONS"
	PerpetualDaylightWaypointTrait = "PERPETUAL_DAYLIGHT"
	PerpetualOvercastWaypointTrait = "PERPETUAL_OVERCAST"
	DrySeabedsWaypointTrait = "DRY_SEABEDS"
	MagmaSeasWaypointTrait = "MAGMA_SEAS"
	SupervolcanoesWaypointTrait = "SUPERVOLCANOES"
	AshCloudsWaypointTrait = "ASH_CLOUDS"
	VastRuinsWaypointTrait = "VAST_RUINS"
	MutatedFloraWaypointTrait = "MUTATED_FLORA"
	TerraformedWaypointTrait = "TERRAFORMED"
	ExtremeTemperaturesWaypointTrait = "EXTREME_TEMPERATURES"
	ExtremePressureWaypointTrait = "EXTREME_PRESSURE"
	DiverseLifeWaypointTrait = "DIVERSE_LIFE"
	ScarceLifeWaypointTrait = "SCARCE_LIFE"
	FossilsWaypointTrait = "FOSSILS"
	WeakGravityWaypointTrait = "WEAK_GRAVITY"
	StrongGravityWaypointTrait = "STRONG_GRAVITY"
	CrushingGravityWaypointTrait = "CRUSHING_GRAVITY"
	ToxicAtmosphereWaypointTrait = "TOXIC_ATMOSPHERE"
	CorrosiveAtmosphereWaypointTrait = "CORROSIVE_ATMOSPHERE"
	BreathableAtmosphereWaypointTrait = "BREATHABLE_ATMOSPHERE"
	ThinAtmosphereWaypointTrait = "THIN_ATMOSPHERE"
	JovianWaypointTrait = "JOVIAN"
	RockyWaypointTrait = "ROCKY"
	VolcanicWaypointTrait = "VOLCANIC"
	FrozenWaypointTrait = "FROZEN"
	SwampWaypointTrait = "SWAMP"
	BarrenWaypointTrait = "BARREN"
	TemperateWaypointTrait = "TEMPERATE"
	JungleWaypointTrait = "JUNGLE"
	OceanWaypointTrait = "OCEAN"
	RadioactiveWaypointTrait = "RADIOACTIVE"
	MicroGravityAnomaliesWaypointTrait = "MICRO_GRAVITY_ANOMALIES"
	DebrisClusterWaypointTrait = "DEBRIS_CLUSTER"
	DeepCratersWaypointTrait = "DEEP_CRATERS"
	ShallowCratersWaypointTrait = "SHALLOW_CRATERS"
	UnstableCompositionWaypointTrait = "UNSTABLE_COMPOSITION"
	HollowedInteriorWaypointTrait = "HOLLOWED_INTERIOR"
	StrippedWaypointTrait = "STRIPPED"
)

type WaypointModifierSymbol string
const (
	StrippedModifier WaypointModifierSymbol = "STRIPPED"
	UnstableModifier = "UNSTABLE"
	RadiationLeakModifier = "RADIATION_LEAK"
	CriticalLimitModifier = "CRITICAL_LIMIT"
	CivilUnrestModifier = "CIVIL_UNREST"
)

type MarketTransactionType string
const (
	PurchaseTransaction MarketTransactionType = "PURCHASE"
	SellTransaction = "SELL"
)

type MarketTradeType string
const (
	ExportTradeType MarketTradeType = "EXPORT"
	ImportTradeType = "IMPORT"
	ExchangeTradeType = "EXCHANGE"
)

type SupplyLevel string
const (
	ScarceSupply SupplyLevel = "SCARCE"
	LimitedSupply = "LIMITED"
	ModerateSupply = "MODERATE"
	HighSupply = "HIGH"
	AbundantSupply = "ABUNDANT"
)

type ActivityLevel string
const (
	UnknownActivity ActivityLevel = ""
	WeakActivity = "WEAK"
	GrowingActivity = "GROWING"
	StrongActivity = "STRONG"
	RestrictedActivity = "RESTRICTED"
)

type SurveySize string
const (
	SmallSize SurveySize = "SMALL"
	ModerateSize = "MODERATE"
	LargeSize = "LARGE"
)
