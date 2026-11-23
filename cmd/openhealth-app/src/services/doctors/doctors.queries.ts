import { queryOptions } from "@tanstack/react-query";

import { GetDoctorByAccountIDFn } from "./doctors.actions";

export const doctorKeys = {
  all: ["doctors"] as const,
  byAccount: (accountID: string) =>
    [...doctorKeys.all, "account", accountID] as const,
};

// GetDoctorByAccountID - Fetches the doctor associated with an account.
export const GetDoctorByAccountID = (accountID: string) =>
  queryOptions({
    queryKey: doctorKeys.byAccount(accountID),
    queryFn: async () => {
      const { doctor, error } = await GetDoctorByAccountIDFn({
        data: { account_id: accountID },
      });

      if (error) {
        throw new Error(error);
      }

      return doctor;
    },
  });

// IsDoctor - Checks whether the given account is a doctor.
export const IsDoctor = (accountID: string) =>
  queryOptions({
    queryKey: [...doctorKeys.byAccount(accountID), "is-doctor"] as const,
    queryFn: async () => {
      const { doctor, error } = await GetDoctorByAccountIDFn({
        data: { account_id: accountID },
      });

      if (error) {
        throw new Error(error);
      }

      return doctor !== null;
    },
  });
