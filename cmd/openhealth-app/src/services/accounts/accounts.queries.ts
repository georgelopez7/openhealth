import { queryOptions } from "@tanstack/react-query";

import { GetAccountsFn } from "./accounts.actions";

export const accountKeys = {
  all: ["accounts"] as const,
  lists: () => [...accountKeys.all, "list"] as const,
  list: (limit: number) => [...accountKeys.lists(), limit] as const,
};

// GetAccounts - Fetches all accounts up to the provided limit.
export const GetAccounts = (limit: number = 1000) =>
  queryOptions({
    queryKey: accountKeys.list(limit),
    queryFn: async () => {
      const { accounts, error } = await GetAccountsFn({ data: { limit } });

      if (error) {
        throw new Error(error);
      }

      return accounts;
    },
  });
