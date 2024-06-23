
class CodeJamEvent {

    Id : string;
    StatusId : number;
    Title : string;
    Description: string;
    Timeline: string;
    Rules: string;

    constructor() {
        this.Id = '';
        this.StatusId = 0;
        this.Title = '';
        this.Description = '';
        this.Rules = '';
        this.Timeline = '';
    }

}

export default CodeJamEvent;