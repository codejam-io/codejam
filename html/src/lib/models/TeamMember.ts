class TeamMember {
    TeamId: string;
    UserId : string;
    Role: string;
    DisplayName: string;
    
    //TODO add Array<Teams> 

    constructor() {
        this.TeamId = '';
        this.UserId = '';
        this.Role = '';
        this.DisplayName = '';
        
    }
}

export default TeamMember;