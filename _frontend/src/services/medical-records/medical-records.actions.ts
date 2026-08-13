import { createServerFn } from "@tanstack/react-start";

import { API_BASE_URL } from "#/domain/config";
import type { MedicalRecord } from "#/domain/medical-records";

export type GetMedicalRecordsResult = {
  medicalRecords: MedicalRecord[];
  error: string | null;
};

// GetMedicalRecordsFn - Fetches all medical records.
export const GetMedicalRecordsFn = createServerFn({ method: "GET" }).handler(
  async (): Promise<GetMedicalRecordsResult> => {
    const response = await fetch(`${API_BASE_URL}/api/v1/medical-records`, {
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      return {
        medicalRecords: [],
        error: `Failed to fetch medical records: ${response.status} ${response.statusText}`,
      };
    }

    type GetMedicalRecordsResponseBody = {
      medical_records: MedicalRecord[] | null;
    };

    const body: GetMedicalRecordsResponseBody = await response.json();
    return {
      medicalRecords: body.medical_records ?? [],
      error: null,
    };
  },
);

export type CreateMedicalRecordResult = {
  medicalRecord: MedicalRecord | null;
  error: string | null;
};

// CreateMedicalRecordFn - Creates a new medical record.
export const CreateMedicalRecordFn = createServerFn({ method: "POST" })
  .validator(
    (data: { account_id: string; title: string; description: string }) => data,
  )
  .handler(async ({ data }): Promise<CreateMedicalRecordResult> => {
    const response = await fetch(`${API_BASE_URL}/api/v1/medical-records`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      return {
        medicalRecord: null,
        error: `Failed to create medical record: ${response.status} ${response.statusText}`,
      };
    }

    type CreateMedicalRecordResponseBody = {
      medical_record: MedicalRecord | null;
    };

    const body: CreateMedicalRecordResponseBody = await response.json();
    return {
      medicalRecord: body.medical_record ?? null,
      error: null,
    };
  });
