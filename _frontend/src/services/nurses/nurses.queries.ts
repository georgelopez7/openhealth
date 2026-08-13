import { queryOptions } from "@tanstack/react-query";

import { GetNurseByAccountIDFn } from "./nurses.actions";

export const nurseKeys = {
  all: ["nurses"] as const,
  byAccount: (accountID: string) =>
    [...nurseKeys.all, "account", accountID] as const,
};

// GetNurseByAccountID - Fetches the nurse associated with an account.
export const GetNurseByAccountID = (accountID: string) =>
  queryOptions({
    queryKey: nurseKeys.byAccount(accountID),
    queryFn: async () => {
      const { nurse, error } = await GetNurseByAccountIDFn({
        data: { account_id: accountID },
      });

      if (error) {
        throw new Error(error);
      }

      return nurse;
    },
  });

// IsNurse - Checks whether the given account is a nurse.
export const IsNurse = (accountID: string) =>
  queryOptions({
    queryKey: [...nurseKeys.byAccount(accountID), "is-nurse"] as const,
    queryFn: async () => {
      const { nurse, error } = await GetNurseByAccountIDFn({
        data: { account_id: accountID },
      });

      if (error) {
        throw new Error(error);
      }

      return nurse !== null;
    },
  });
