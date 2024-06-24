export interface User {
    UserId: string;
    DisplayName: string;
    Role: string;
    ServiceName : string;
    ServiceUserId: string;
    AvatarUrl: string;
}

export interface ActiveUser {
    user : User | null;
    loggedIn : boolean;
}