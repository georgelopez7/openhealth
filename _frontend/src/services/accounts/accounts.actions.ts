import { createServerFn } from "@tanstack/react-start";

import type { Account } from "#/domain/accounts";
import { API_BASE_URL } from "#/domain/config";

export type GetAccountsResult = {
  accounts: Account[];
  error: string | null;
};

// GetAccountsFn - Fetches all accounts up to the provided limit.
export const GetAccountsFn = createServerFn({ method: "GET" })
  .validator((data: { limit?: number }) => data)
  .handler(async ({ data }): Promise<GetAccountsResult> => {
    const url = new URL("/api/v1/accounts", API_BASE_URL);
    url.searchParams.set("limit", String(data.limit ?? 1000));

    const response = await fetch(url.toString(), {
      headers: {
        Accept: "application/json",
      },
    });

    if (!response.ok) {
      return {
        accounts: [],
        error: `Failed to fetch accounts: ${response.status} ${response.statusText}`,
      };
    }

    type GetAccountsResponseBody = {
      accounts: Account[] | null;
    };

    const body: GetAccountsResponseBody = await response.json();
    return {
      accounts: body.accounts ?? [],
      error: null,
    };
  });

export type CreateAccountResult = {
  account: Account | null;
  error: string | null;
};

// CreateAccountFn - Creates a new account.
export const CreateAccountFn = createServerFn({ method: "POST" })
  .validator(
    (data: {
      first_name: string;
      last_name: string;
      age: number;
      email: string;
    }) => data,
  )
  .handler(async ({ data }): Promise<CreateAccountResult> => {
    const response = await fetch(`${API_BASE_URL}/api/v1/accounts`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      return {
        account: null,
        error: `Failed to create account: ${response.status} ${response.statusText}`,
      };
    }

    type CreateAccountResponseBody = {
      account: Account | null;
    };

    const body: CreateAccountResponseBody = await response.json();
    return {
      account: body.account ?? null,
      error: null,
    };
  });
