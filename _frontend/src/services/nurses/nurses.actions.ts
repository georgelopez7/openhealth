import { createServerFn } from "@tanstack/react-start";

import { API_BASE_URL } from "#/domain/config";
import type { Nurse } from "#/domain/nurses";

export type GetNurseByAccountIDResult = {
  nurse: Nurse | null;
  error: string | null;
};

// GetNurseByAccountIDFn - Fetches the nurse associated with an account.
export const GetNurseByAccountIDFn = createServerFn({ method: "GET" })
  .validator((data: { account_id: string }) => data)
  .handler(async ({ data }): Promise<GetNurseByAccountIDResult> => {
    const response = await fetch(
      `${API_BASE_URL}/api/v1/accounts/${data.account_id}/nurse`,
      {
        headers: {
          Accept: "application/json",
        },
      },
    );

    if (response.status === 404) {
      return { nurse: null, error: null };
    }

    if (!response.ok) {
      return {
        nurse: null,
        error: `Failed to fetch nurse: ${response.status} ${response.statusText}`,
      };
    }

    type GetNurseByAccountIDResponseBody = {
      nurse: Nurse | null;
    };

    const body: GetNurseByAccountIDResponseBody = await response.json();
    return {
      nurse: body.nurse ?? null,
      error: null,
    };
  });
