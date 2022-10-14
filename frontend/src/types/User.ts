export interface User {
    id: number;
    name: string;
    lastName: string;
}

export interface Profile {
    id: number;
    name: string;
    tag: string;
}

export type Selection = 'profile' | 'user';
