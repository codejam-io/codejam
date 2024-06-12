class TeamMember {
    TeamId: string;
    UserId : string;
    TeamRole: string;
    DisplayName: string;
    
    //TODO add Array<Teams> 

    constructor() {
        this.TeamId = '';
        this.UserId = '';
        this.TeamRole = '';
        this.DisplayName = '';
        
    }
}

export default TeamMember;