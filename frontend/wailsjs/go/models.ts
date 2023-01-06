export namespace fireapi {
	
	export class Report {
	    region: string;
	    reason: string;
	    level: string;
	    reportType: number;
	    extraGiftId: number;
	    reportedName: string;
	    creatorName: string;
	    extInfo: string;
	
	    static createFrom(source: any = {}) {
	        return new Report(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.region = source["region"];
	        this.reason = source["reason"];
	        this.level = source["level"];
	        this.reportType = source["reportType"];
	        this.extraGiftId = source["extraGiftId"];
	        this.reportedName = source["reportedName"];
	        this.creatorName = source["creatorName"];
	        this.extInfo = source["extInfo"];
	    }
	}

}

export namespace lcu {
	
	export class AdditionalProp1 {
	    spells: number[];
	
	    static createFrom(source: any = {}) {
	        return new AdditionalProp1(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.spells = source["spells"];
	    }
	}
	export class Bans {
	    championId: number;
	    pickTurn: number;
	
	    static createFrom(source: any = {}) {
	        return new Bans(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.championId = source["championId"];
	        this.pickTurn = source["pickTurn"];
	    }
	}
	export class CreepsPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new CreepsPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class CsDiffPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new CsDiffPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class DamageTakenDiffPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new DamageTakenDiffPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class DamageTakenPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new DamageTakenPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class Teams {
	    bans: Bans[];
	    baronKills: number;
	    dominionVictoryScore: number;
	    dragonKills: number;
	    firstBaron: boolean;
	    firstBlood: boolean;
	    firstDargon: boolean;
	    firstInhibitor: boolean;
	    firstTower: boolean;
	    inhibitorKills: number;
	    riftHeraldKills: number;
	    teamId: number;
	    towerKills: number;
	    vilemawKills: number;
	    win: string;
	
	    static createFrom(source: any = {}) {
	        return new Teams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bans = this.convertValues(source["bans"], Bans);
	        this.baronKills = source["baronKills"];
	        this.dominionVictoryScore = source["dominionVictoryScore"];
	        this.dragonKills = source["dragonKills"];
	        this.firstBaron = source["firstBaron"];
	        this.firstBlood = source["firstBlood"];
	        this.firstDargon = source["firstDargon"];
	        this.firstInhibitor = source["firstInhibitor"];
	        this.firstTower = source["firstTower"];
	        this.inhibitorKills = source["inhibitorKills"];
	        this.riftHeraldKills = source["riftHeraldKills"];
	        this.teamId = source["teamId"];
	        this.towerKills = source["towerKills"];
	        this.vilemawKills = source["vilemawKills"];
	        this.win = source["win"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class XpPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new XpPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class XpDiffPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new XpDiffPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class GoldPerMinDeltas {
	    "0-10": number;
	    "10-20": number;
	    "20-30": number;
	    "30-end": number;
	
	    static createFrom(source: any = {}) {
	        return new GoldPerMinDeltas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["0-10"] = source["0-10"];
	        this["10-20"] = source["10-20"];
	        this["20-30"] = source["20-30"];
	        this["30-end"] = source["30-end"];
	    }
	}
	export class Timeline {
	    creepsPerMinDeltas: CreepsPerMinDeltas;
	    csDiffPerMinDeltas: CsDiffPerMinDeltas;
	    damageTakenDiffPerMinDeltas: DamageTakenDiffPerMinDeltas;
	    damageTakenPerMinDeltas: DamageTakenPerMinDeltas;
	    goldPerMinDeltas: GoldPerMinDeltas;
	    lane: string;
	    participantId: number;
	    role: string;
	    xpDiffPerMinDeltas: XpDiffPerMinDeltas;
	    xpPerMinDeltas: XpPerMinDeltas;
	
	    static createFrom(source: any = {}) {
	        return new Timeline(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.creepsPerMinDeltas = this.convertValues(source["creepsPerMinDeltas"], CreepsPerMinDeltas);
	        this.csDiffPerMinDeltas = this.convertValues(source["csDiffPerMinDeltas"], CsDiffPerMinDeltas);
	        this.damageTakenDiffPerMinDeltas = this.convertValues(source["damageTakenDiffPerMinDeltas"], DamageTakenDiffPerMinDeltas);
	        this.damageTakenPerMinDeltas = this.convertValues(source["damageTakenPerMinDeltas"], DamageTakenPerMinDeltas);
	        this.goldPerMinDeltas = this.convertValues(source["goldPerMinDeltas"], GoldPerMinDeltas);
	        this.lane = source["lane"];
	        this.participantId = source["participantId"];
	        this.role = source["role"];
	        this.xpDiffPerMinDeltas = this.convertValues(source["xpDiffPerMinDeltas"], XpDiffPerMinDeltas);
	        this.xpPerMinDeltas = this.convertValues(source["xpPerMinDeltas"], XpPerMinDeltas);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Participants {
	    championId: number;
	    highestAchievedSeasonTier: string;
	    participantId: number;
	    spell1Id: number;
	    spell2Id: number;
	    stats: Stats;
	    teamId: number;
	    timeline: Timeline;
	
	    static createFrom(source: any = {}) {
	        return new Participants(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.championId = source["championId"];
	        this.highestAchievedSeasonTier = source["highestAchievedSeasonTier"];
	        this.participantId = source["participantId"];
	        this.spell1Id = source["spell1Id"];
	        this.spell2Id = source["spell2Id"];
	        this.stats = this.convertValues(source["stats"], Stats);
	        this.teamId = source["teamId"];
	        this.timeline = this.convertValues(source["timeline"], Timeline);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Stats {
	    assists: number;
	    causedEarlySurrender: boolean;
	    champLevel: number;
	    combatPlayerScore: number;
	    damageDealtToObjectives: number;
	    damageDealtToTurrets: number;
	    damageSelfMitigated: number;
	    deaths: number;
	    doubleKills: number;
	    earlySurrenderAccomplice: boolean;
	    firstBloodAssist: boolean;
	    firstBloodKill: boolean;
	    firstInhibitorAssist: boolean;
	    firstInhibitorKill: boolean;
	    firstTowerAssist: boolean;
	    firstTowerKill: boolean;
	    gameEndedInEarlySurrender: boolean;
	    gameEndedInSurrender: boolean;
	    goldEarned: number;
	    goldSpent: number;
	    inhibitorKills: number;
	    item0: number;
	    item1: number;
	    item2: number;
	    item3: number;
	    item4: number;
	    item5: number;
	    item6: number;
	    killingSprees: number;
	    kills: number;
	    largestCriticalStrike: number;
	    largestKillingSpree: number;
	    largestMultiKill: number;
	    longestTimeSpentLiving: number;
	    magicDamageDealt: number;
	    magicDamageDealtToChampions: number;
	    magicalDamageTaken: number;
	    neutralMinionsKilled: number;
	    neutralMinionsKilledEnemyJungle: number;
	    neutralMinionsKilledTeamJungle: number;
	    objectivePlayerScore: number;
	    participantId: number;
	    pentaKills: number;
	    perk0: number;
	    perk0Var1: number;
	    perk0Var2: number;
	    perk0Var3: number;
	    perk1: number;
	    perk1Var1: number;
	    perk1Var2: number;
	    perk1Var3: number;
	    perk2: number;
	    perk2Var1: number;
	    perk2Var2: number;
	    perk2Var3: number;
	    perk3: number;
	    perk3Var1: number;
	    perk3Var2: number;
	    perk3Var3: number;
	    perk4: number;
	    perk4Var1: number;
	    perk4Var2: number;
	    perk4Var3: number;
	    perk5: number;
	    perk5Var1: number;
	    perk5Var2: number;
	    perk5Var3: number;
	    perkPrimaryStyle: number;
	    perkSubStyle: number;
	    physicalDamageDealt: number;
	    physicalDamageDealtToChampions: number;
	    physicalDamageTaken: number;
	    playerScore0: number;
	    playerScore1: number;
	    playerScore2: number;
	    playerScore3: number;
	    playerScore4: number;
	    playerScore5: number;
	    playerScore6: number;
	    playerScore7: number;
	    playerScore8: number;
	    playerScore9: number;
	    quadraKills: number;
	    sightWardsBoughtInGame: number;
	    teamEarlySurrendered: boolean;
	    timeCCingOthers: number;
	    totalDamageDealt: number;
	    totalDamageDealtToChampions: number;
	    totalDamageTaken: number;
	    totalHeal: number;
	    totalMinionsKilled: number;
	    totalPlayerScore: number;
	    totalScoreRank: number;
	    totalTimeCrowdControlDealt: number;
	    totalUnitsHealed: number;
	    tripleKills: number;
	    trueDamageDealt: number;
	    trueDamageDealtToChampions: number;
	    trueDamageTaken: number;
	    turretKills: number;
	    unrealKills: number;
	    visionScore: number;
	    visionWardsBoughtInGame: number;
	    wardsKilled: number;
	    wardsPlaced: number;
	    win: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Stats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assists = source["assists"];
	        this.causedEarlySurrender = source["causedEarlySurrender"];
	        this.champLevel = source["champLevel"];
	        this.combatPlayerScore = source["combatPlayerScore"];
	        this.damageDealtToObjectives = source["damageDealtToObjectives"];
	        this.damageDealtToTurrets = source["damageDealtToTurrets"];
	        this.damageSelfMitigated = source["damageSelfMitigated"];
	        this.deaths = source["deaths"];
	        this.doubleKills = source["doubleKills"];
	        this.earlySurrenderAccomplice = source["earlySurrenderAccomplice"];
	        this.firstBloodAssist = source["firstBloodAssist"];
	        this.firstBloodKill = source["firstBloodKill"];
	        this.firstInhibitorAssist = source["firstInhibitorAssist"];
	        this.firstInhibitorKill = source["firstInhibitorKill"];
	        this.firstTowerAssist = source["firstTowerAssist"];
	        this.firstTowerKill = source["firstTowerKill"];
	        this.gameEndedInEarlySurrender = source["gameEndedInEarlySurrender"];
	        this.gameEndedInSurrender = source["gameEndedInSurrender"];
	        this.goldEarned = source["goldEarned"];
	        this.goldSpent = source["goldSpent"];
	        this.inhibitorKills = source["inhibitorKills"];
	        this.item0 = source["item0"];
	        this.item1 = source["item1"];
	        this.item2 = source["item2"];
	        this.item3 = source["item3"];
	        this.item4 = source["item4"];
	        this.item5 = source["item5"];
	        this.item6 = source["item6"];
	        this.killingSprees = source["killingSprees"];
	        this.kills = source["kills"];
	        this.largestCriticalStrike = source["largestCriticalStrike"];
	        this.largestKillingSpree = source["largestKillingSpree"];
	        this.largestMultiKill = source["largestMultiKill"];
	        this.longestTimeSpentLiving = source["longestTimeSpentLiving"];
	        this.magicDamageDealt = source["magicDamageDealt"];
	        this.magicDamageDealtToChampions = source["magicDamageDealtToChampions"];
	        this.magicalDamageTaken = source["magicalDamageTaken"];
	        this.neutralMinionsKilled = source["neutralMinionsKilled"];
	        this.neutralMinionsKilledEnemyJungle = source["neutralMinionsKilledEnemyJungle"];
	        this.neutralMinionsKilledTeamJungle = source["neutralMinionsKilledTeamJungle"];
	        this.objectivePlayerScore = source["objectivePlayerScore"];
	        this.participantId = source["participantId"];
	        this.pentaKills = source["pentaKills"];
	        this.perk0 = source["perk0"];
	        this.perk0Var1 = source["perk0Var1"];
	        this.perk0Var2 = source["perk0Var2"];
	        this.perk0Var3 = source["perk0Var3"];
	        this.perk1 = source["perk1"];
	        this.perk1Var1 = source["perk1Var1"];
	        this.perk1Var2 = source["perk1Var2"];
	        this.perk1Var3 = source["perk1Var3"];
	        this.perk2 = source["perk2"];
	        this.perk2Var1 = source["perk2Var1"];
	        this.perk2Var2 = source["perk2Var2"];
	        this.perk2Var3 = source["perk2Var3"];
	        this.perk3 = source["perk3"];
	        this.perk3Var1 = source["perk3Var1"];
	        this.perk3Var2 = source["perk3Var2"];
	        this.perk3Var3 = source["perk3Var3"];
	        this.perk4 = source["perk4"];
	        this.perk4Var1 = source["perk4Var1"];
	        this.perk4Var2 = source["perk4Var2"];
	        this.perk4Var3 = source["perk4Var3"];
	        this.perk5 = source["perk5"];
	        this.perk5Var1 = source["perk5Var1"];
	        this.perk5Var2 = source["perk5Var2"];
	        this.perk5Var3 = source["perk5Var3"];
	        this.perkPrimaryStyle = source["perkPrimaryStyle"];
	        this.perkSubStyle = source["perkSubStyle"];
	        this.physicalDamageDealt = source["physicalDamageDealt"];
	        this.physicalDamageDealtToChampions = source["physicalDamageDealtToChampions"];
	        this.physicalDamageTaken = source["physicalDamageTaken"];
	        this.playerScore0 = source["playerScore0"];
	        this.playerScore1 = source["playerScore1"];
	        this.playerScore2 = source["playerScore2"];
	        this.playerScore3 = source["playerScore3"];
	        this.playerScore4 = source["playerScore4"];
	        this.playerScore5 = source["playerScore5"];
	        this.playerScore6 = source["playerScore6"];
	        this.playerScore7 = source["playerScore7"];
	        this.playerScore8 = source["playerScore8"];
	        this.playerScore9 = source["playerScore9"];
	        this.quadraKills = source["quadraKills"];
	        this.sightWardsBoughtInGame = source["sightWardsBoughtInGame"];
	        this.teamEarlySurrendered = source["teamEarlySurrendered"];
	        this.timeCCingOthers = source["timeCCingOthers"];
	        this.totalDamageDealt = source["totalDamageDealt"];
	        this.totalDamageDealtToChampions = source["totalDamageDealtToChampions"];
	        this.totalDamageTaken = source["totalDamageTaken"];
	        this.totalHeal = source["totalHeal"];
	        this.totalMinionsKilled = source["totalMinionsKilled"];
	        this.totalPlayerScore = source["totalPlayerScore"];
	        this.totalScoreRank = source["totalScoreRank"];
	        this.totalTimeCrowdControlDealt = source["totalTimeCrowdControlDealt"];
	        this.totalUnitsHealed = source["totalUnitsHealed"];
	        this.tripleKills = source["tripleKills"];
	        this.trueDamageDealt = source["trueDamageDealt"];
	        this.trueDamageDealtToChampions = source["trueDamageDealtToChampions"];
	        this.trueDamageTaken = source["trueDamageTaken"];
	        this.turretKills = source["turretKills"];
	        this.unrealKills = source["unrealKills"];
	        this.visionScore = source["visionScore"];
	        this.visionWardsBoughtInGame = source["visionWardsBoughtInGame"];
	        this.wardsKilled = source["wardsKilled"];
	        this.wardsPlaced = source["wardsPlaced"];
	        this.win = source["win"];
	    }
	}
	export class Player {
	    accountId: number;
	    currentAccountId: number;
	    currentPlatformId: string;
	    matchHistoryUri: string;
	    platformId: string;
	    profileIcon: number;
	    summonerId: number;
	    summonerName: string;
	
	    static createFrom(source: any = {}) {
	        return new Player(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.currentAccountId = source["currentAccountId"];
	        this.currentPlatformId = source["currentPlatformId"];
	        this.matchHistoryUri = source["matchHistoryUri"];
	        this.platformId = source["platformId"];
	        this.profileIcon = source["profileIcon"];
	        this.summonerId = source["summonerId"];
	        this.summonerName = source["summonerName"];
	    }
	}
	export class ParticipantIdentities {
	    participantId: number;
	    player: Player;
	    stats?: Stats;
	
	    static createFrom(source: any = {}) {
	        return new ParticipantIdentities(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.participantId = source["participantId"];
	        this.player = this.convertValues(source["player"], Player);
	        this.stats = this.convertValues(source["stats"], Stats);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MatchInfo {
	    gameCreation: number;
	    gameCreationDate: string;
	    gameDuration: number;
	    gameId: number;
	    gameMode: string;
	    gameType: string;
	    gameVersion: string;
	    mapId: number;
	    participantIdentities: ParticipantIdentities[];
	    participants: Participants[];
	    platformId: string;
	    queueId: number;
	    seasonId: number;
	    teams: Teams[];
	
	    static createFrom(source: any = {}) {
	        return new MatchInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.gameCreation = source["gameCreation"];
	        this.gameCreationDate = source["gameCreationDate"];
	        this.gameDuration = source["gameDuration"];
	        this.gameId = source["gameId"];
	        this.gameMode = source["gameMode"];
	        this.gameType = source["gameType"];
	        this.gameVersion = source["gameVersion"];
	        this.mapId = source["mapId"];
	        this.participantIdentities = this.convertValues(source["participantIdentities"], ParticipantIdentities);
	        this.participants = this.convertValues(source["participants"], Participants);
	        this.platformId = source["platformId"];
	        this.queueId = source["queueId"];
	        this.seasonId = source["seasonId"];
	        this.teams = this.convertValues(source["teams"], Teams);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Games {
	    gameBeginDate: string;
	    gameCount: number;
	    gameEndDate: string;
	    gameIndexBegin: number;
	    gameIndexEnd: number;
	    games: MatchInfo[];
	
	    static createFrom(source: any = {}) {
	        return new Games(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.gameBeginDate = source["gameBeginDate"];
	        this.gameCount = source["gameCount"];
	        this.gameEndDate = source["gameEndDate"];
	        this.gameIndexBegin = source["gameIndexBegin"];
	        this.gameIndexEnd = source["gameIndexEnd"];
	        this.games = this.convertValues(source["games"], MatchInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class MatchList {
	    accountId: number;
	    games: Games;
	    platformId: string;
	
	    static createFrom(source: any = {}) {
	        return new MatchList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.games = this.convertValues(source["games"], Games);
	        this.platformId = source["platformId"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	export class RerollPoints {
	    currentPoints: number;
	    maxRolls: number;
	    numberOfRolls: number;
	    pointsCostToRoll: number;
	    pointsToReroll: number;
	
	    static createFrom(source: any = {}) {
	        return new RerollPoints(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.currentPoints = source["currentPoints"];
	        this.maxRolls = source["maxRolls"];
	        this.numberOfRolls = source["numberOfRolls"];
	        this.pointsCostToRoll = source["pointsCostToRoll"];
	        this.pointsToReroll = source["pointsToReroll"];
	    }
	}
	
	export class Summoner {
	    accountId: number;
	    displayName: string;
	    internalName: string;
	    nameChangeFlag: boolean;
	    percentCompleteForNextLevel: number;
	    privacy: string;
	    profileIconId: number;
	    puuid: string;
	    rerollPoints: RerollPoints;
	    summonerId: number;
	    summonerLevel: number;
	    unnamed: boolean;
	    xpSinceLastLevel: number;
	    xpUntilNextLevel: number;
	
	    static createFrom(source: any = {}) {
	        return new Summoner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accountId = source["accountId"];
	        this.displayName = source["displayName"];
	        this.internalName = source["internalName"];
	        this.nameChangeFlag = source["nameChangeFlag"];
	        this.percentCompleteForNextLevel = source["percentCompleteForNextLevel"];
	        this.privacy = source["privacy"];
	        this.profileIconId = source["profileIconId"];
	        this.puuid = source["puuid"];
	        this.rerollPoints = this.convertValues(source["rerollPoints"], RerollPoints);
	        this.summonerId = source["summonerId"];
	        this.summonerLevel = source["summonerLevel"];
	        this.unnamed = source["unnamed"];
	        this.xpSinceLastLevel = source["xpSinceLastLevel"];
	        this.xpUntilNextLevel = source["xpUntilNextLevel"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Team {
	    additionalProp1: AdditionalProp1;
	    accountId: number;
	    summonerName: string;
	    summonerInternalName: string;
	    selectedPosition: string;
	    championId: number;
	    positionId: number;
	
	    static createFrom(source: any = {}) {
	        return new Team(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.additionalProp1 = this.convertValues(source["additionalProp1"], AdditionalProp1);
	        this.accountId = source["accountId"];
	        this.summonerName = source["summonerName"];
	        this.summonerInternalName = source["summonerInternalName"];
	        this.selectedPosition = source["selectedPosition"];
	        this.championId = source["championId"];
	        this.positionId = source["positionId"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SummonerListInGame {
	    friendTeam: Team[];
	    enemyTeam: Team[];
	
	    static createFrom(source: any = {}) {
	        return new SummonerListInGame(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.friendTeam = this.convertValues(source["friendTeam"], Team);
	        this.enemyTeam = this.convertValues(source["enemyTeam"], Team);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	

}

