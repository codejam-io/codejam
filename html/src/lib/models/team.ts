import TeamMember from "./TeamMember";

class CodeJamTeam {

    Id : string;
    EventId: string;
    Name: string;
    Visibility: string;
    Technologies: string;
    Availability: string;
    Description: string;
    InviteCode: string;

    constructor() {
        this.Id = '';
        this.EventId = '';
        this.Name = '';
        this.Visibility = '';
        this.Technologies = '';
        this.Availability = '';
        this.Description = '';
        this.InviteCode = '';
    }
}

export default CodeJamTeam;