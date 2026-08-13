export enum MedicalRecordStatus {
  Active = "active",
  Archived = "archived",
}

export interface MedicalRecord {
  id: string;
  account_id: string;
  title: string;
  description: string;
  status: MedicalRecordStatus;
  created_at: string;
  updated_at: string;
}
