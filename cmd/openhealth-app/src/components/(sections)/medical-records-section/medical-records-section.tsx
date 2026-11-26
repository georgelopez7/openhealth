import { useQueries } from "@tanstack/react-query";
import { Eye, LoaderCircle, Pencil } from "lucide-react";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "#/components/ui/tooltip";
import type { Account } from "#/domain/accounts";
import { getAccountDisplayName } from "#/domain/accounts";
import { AccountIcon } from "#/domain/icon";
import type { MedicalRecord } from "#/domain/medical-records";
import { cn } from "#/lib/utils";
import { GetMedicalRecordAccess } from "#/services/medical-records/medical-records.queries";

interface IPermissionIconProps {
  action: "view" | "edit";
  granted: boolean;
  isLoading: boolean;
}

const PermissionIcon = ({
  action,
  granted,
  isLoading,
}: IPermissionIconProps) => {
  const Icon = isLoading ? LoaderCircle : action === "view" ? Eye : Pencil;
  const label = isLoading
    ? `Checking ${action} access`
    : granted
      ? `Can ${action}`
      : `Cannot ${action}`;

  return (
    <Tooltip>
      <TooltipTrigger
        type="button"
        aria-label={label}
        className={cn(
          "inline-flex size-8 items-center justify-center border border-neutral-700 bg-[#101010] transition-colors hover:border-neutral-400 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-white/50",
          isLoading
            ? "cursor-wait text-neutral-500"
            : granted
              ? "text-white"
              : "text-white/25",
        )}
      >
        <Icon
          aria-hidden="true"
          className={cn("size-4", isLoading && "animate-spin")}
        />
      </TooltipTrigger>
      <TooltipContent className="border border-neutral-700 bg-[#101010] text-xs text-white">
        {label}
      </TooltipContent>
    </Tooltip>
  );
};

interface IProps {
  account: Account;
  accounts: Account[];
  medicalRecords: MedicalRecord[];
}

const MedicalRecordsSection = ({
  account,
  accounts,
  medicalRecords,
}: IProps) => {
  const accountsByID = new Map(
    accounts.map((candidate) => [candidate.id, candidate]),
  );
  const accessResults = useQueries({
    queries: medicalRecords.map((record) =>
      GetMedicalRecordAccess(record.id, account.id),
    ),
    combine: (results) =>
      results.map((result) => ({
        canView: result.data?.can_view ?? false,
        canEdit: result.data?.can_edit ?? false,
        isLoading: result.isLoading,
      })),
  });

  return (
    <section className="min-w-0 flex-1 text-white">
      <div className="w-full overflow-x-auto border border-neutral-700">
        <table className="w-full min-w-[44rem] table-fixed text-left">
          <caption className="sr-only">
            Medical records and the selected character&apos;s access
          </caption>
          <colgroup>
            <col />
            <col className="w-40" />
            <col className="w-20" />
            <col className="w-20" />
          </colgroup>
          <thead className="bg-[#101010]">
            <tr className="border-b border-neutral-700">
              <th
                scope="col"
                className="px-4 py-3 text-[10px] font-semibold uppercase tracking-wider text-neutral-400"
              >
                Record
              </th>
              <th
                scope="col"
                className="px-3 py-3 text-center text-[10px] font-semibold uppercase tracking-wider text-neutral-400"
              >
                Owner
              </th>
              <th
                scope="col"
                className="px-2 py-3 text-center text-[10px] font-semibold uppercase tracking-wider text-neutral-400"
              >
                View
              </th>
              <th
                scope="col"
                className="px-2 py-3 text-center text-[10px] font-semibold uppercase tracking-wider text-neutral-400"
              >
                Edit
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-neutral-800">
            {medicalRecords.length === 0 ? (
              <tr>
                <td
                  colSpan={4}
                  className="px-4 py-8 text-center text-xs text-neutral-500"
                >
                  No medical records found.
                </td>
              </tr>
            ) : (
              medicalRecords.map((record, index) => {
                const access = accessResults[index];
                const owner = accountsByID.get(record.account_id);

                return (
                  <tr key={record.id} className="hover:bg-white/[0.02]">
                    <td className="max-w-0 px-4 py-3">
                      <p className="truncate text-sm font-semibold text-white">
                        {record.title}
                      </p>
                      <p
                        className="mt-1 truncate text-[10px] text-neutral-500"
                        title={record.description}
                      >
                        {record.description}
                      </p>
                    </td>
                    <td
                      aria-label={
                        owner
                          ? `Owner: ${getAccountDisplayName(owner)}`
                          : "Owner unknown"
                      }
                      className="px-3 py-3 text-center"
                    >
                      {owner ? (
                        <AccountIcon
                          avatar={owner.avatar}
                          className="mx-auto size-8"
                        />
                      ) : (
                        <span className="text-neutral-600">—</span>
                      )}
                    </td>
                    <td className="px-2 py-2 text-center">
                      <PermissionIcon
                        action="view"
                        granted={access?.canView ?? false}
                        isLoading={access?.isLoading ?? true}
                      />
                    </td>
                    <td className="px-2 py-2 text-center">
                      <PermissionIcon
                        action="edit"
                        granted={access?.canEdit ?? false}
                        isLoading={access?.isLoading ?? true}
                      />
                    </td>
                  </tr>
                );
              })
            )}
          </tbody>
        </table>
      </div>
    </section>
  );
};

export default MedicalRecordsSection;
