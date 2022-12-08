export namespace lcu {
	
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

}

