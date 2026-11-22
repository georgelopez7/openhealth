export enum NurseStatus {
  Active = "active",
  Archived = "archived",
}

export interface Nurse {
  id: string;
  account_id: string;
  status: NurseStatus;
  created_at: string;
}
