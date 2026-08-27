import { createServerFn } from "@tanstack/react-start";

import { API_AUTH_TOKEN, API_URL } from "#/domain/config";
import type { Doctor } from "#/domain/doctors";

export type GetDoctorByAccountIDResult = {
  doctor: Doctor | null;
  error: string | null;
};

// GetDoctorByAccountIDFn - Fetches the doctor associated with an account.
export const GetDoctorByAccountIDFn = createServerFn({ method: "GET" })
  .validator((data: { account_id: string }) => data)
  .handler(async ({ data }): Promise<GetDoctorByAccountIDResult> => {
    const response = await fetch(
      `${API_URL}/api/v1/accounts/${data.account_id}/doctor`,
      {
        headers: {
          Accept: "application/json",
          Authorization: `Bearer ${API_AUTH_TOKEN}`,
        },
      },
    );

    if (response.status === 404) {
      return { doctor: null, error: null };
    }

    if (!response.ok) {
      return {
        doctor: null,
        error: `Failed to fetch doctor: ${response.status} ${response.statusText}`,
      };
    }

    type GetDoctorByAccountIDResponseBody = {
      doctor: Doctor | null;
    };

    const body: GetDoctorByAccountIDResponseBody = await response.json();
    return {
      doctor: body.doctor ?? null,
      error: null,
    };
  });
