export enum AccountStatus {
  Active = "active",
  Archived = "archived",
}

export interface Account {
  id: string;
  first_name: string;
  last_name: string;
  age: number;
  email: string;
  status: AccountStatus;
  doctor_id?: string;
  hospital_id?: string;
  created_at: string;
  updated_at: string;
}
