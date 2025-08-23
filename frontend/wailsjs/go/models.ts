export namespace main {
	
	export class Repository {
	    repositoryId: string;
	    repositoryUrl: string;
	    username: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Repository(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.repositoryId = source["repositoryId"];
	        this.repositoryUrl = source["repositoryUrl"];
	        this.username = source["username"];
	        this.password = source["password"];
	    }
	}
	export class UploadFileData {
	    artifactId: string;
	    groupId: string;
	    version: string;
	    uploadFileName: string;
	    base64FileData: string;
	    uploadFilePath: string;
	
	    static createFrom(source: any = {}) {
	        return new UploadFileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.artifactId = source["artifactId"];
	        this.groupId = source["groupId"];
	        this.version = source["version"];
	        this.uploadFileName = source["uploadFileName"];
	        this.base64FileData = source["base64FileData"];
	        this.uploadFilePath = source["uploadFilePath"];
	    }
	}

}

