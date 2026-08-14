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

export type GetMedicalRecordAccessResult = {
  access: {
    can_view: boolean;
    can_edit: boolean;
  } | null;
  error: string | null;
};

// GetMedicalRecordAccessFn - Fetches access permissions for a medical record.
export const GetMedicalRecordAccessFn = createServerFn({ method: "GET" })
  .validator((data: { record_id: string; account_id: string }) => data)
  .handler(async ({ data }): Promise<GetMedicalRecordAccessResult> => {
    const response = await fetch(
      `${API_BASE_URL}/api/v1/accounts/${data.account_id}/medical-records/${data.record_id}/access`,
      {
        headers: {
          Accept: "application/json",
        },
      },
    );

    if (!response.ok) {
      return {
        access: null,
        error: `Failed to fetch medical record access: ${response.status} ${response.statusText}`,
      };
    }

    type GetMedicalRecordAccessResponseBody = {
      can_view: boolean;
      can_edit: boolean;
    };

    const body: GetMedicalRecordAccessResponseBody = await response.json();
    return {
      access: body,
      error: null,
    };
  });

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
