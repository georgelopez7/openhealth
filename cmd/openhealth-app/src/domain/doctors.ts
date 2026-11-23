export enum DoctorStatus {
  Active = "active",
  Archived = "archived",
}

export interface Doctor {
  id: string;
  account_id: string;
  status: DoctorStatus;
  created_at: string;
}
