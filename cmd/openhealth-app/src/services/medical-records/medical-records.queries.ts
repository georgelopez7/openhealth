import { queryOptions } from "@tanstack/react-query";

import {
  GetMedicalRecordAccessFn,
  GetMedicalRecordsFn,
} from "./medical-records.actions";

export const medicalRecordKeys = {
  all: ["medical-records"] as const,
  lists: () => [...medicalRecordKeys.all, "list"] as const,
  access: (recordID: string, accountID: string) =>
    [...medicalRecordKeys.all, "access", recordID, accountID] as const,
};

// GetMedicalRecords - Fetches all medical records.
export const GetMedicalRecords = () =>
  queryOptions({
    queryKey: medicalRecordKeys.lists(),
    queryFn: async () => {
      const { medicalRecords, error } = await GetMedicalRecordsFn();

      if (error) {
        throw new Error(error);
      }

      return medicalRecords;
    },
  });

// GetMedicalRecordAccess - Fetches view/edit access for a medical record.
export const GetMedicalRecordAccess = (recordID: string, accountID: string) =>
  queryOptions({
    queryKey: medicalRecordKeys.access(recordID, accountID),
    queryFn: async () => {
      const { access, error } = await GetMedicalRecordAccessFn({
        data: { record_id: recordID, account_id: accountID },
      });

      if (error) {
        throw new Error(error);
      }

      return access;
    },
    enabled: Boolean(recordID) && Boolean(accountID),
  });
