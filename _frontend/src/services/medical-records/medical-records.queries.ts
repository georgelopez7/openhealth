import { queryOptions } from "@tanstack/react-query";

import { GetMedicalRecordsFn } from "./medical-records.actions";

export const medicalRecordKeys = {
  all: ["medical-records"] as const,
  lists: () => [...medicalRecordKeys.all, "list"] as const,
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
