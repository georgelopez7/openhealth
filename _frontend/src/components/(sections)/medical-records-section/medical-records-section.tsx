import { useQueries } from "@tanstack/react-query";
import { Eye, EyeOff, Pencil } from "lucide-react";

import Spacer from "#/components/(layouts)/spacer/spacer";
import type { Account } from "#/domain/accounts";
import type { MedicalRecord } from "#/domain/medical-records";
import { GetMedicalRecordAccess } from "#/services/medical-records/medical-records.queries";

interface IProps {
  account: Account;
  accounts: Account[];
  medicalRecords: MedicalRecord[];
}

// getOwner - Returns the owner's name or the account ID if the owner is not found.
const getOwner = (accounts: Account[], accountID: string): string => {
  const owner = accounts.find((account) => account.id === accountID);
  return owner ? `${owner.first_name} ${owner.last_name}` : accountID;
};

const MedicalRecordsSection = ({
  account,
  accounts,
  medicalRecords,
}: IProps) => {
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

  if (medicalRecords.length === 0) {
    return <p className="opacity-80">No medical records found.</p>;
  }

  return (
    <section className="w-full space-y-4 text-white">
      <h2 className="text-xl font-bold">Medical Records</h2>
      <ul className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {medicalRecords.map((record, index) => {
          const { canView, canEdit, isLoading } = accessResults[index];

          return (
            <li
              key={record.id}
              className="flex min-h-45 flex-col justify-between rounded-lg border border-white/20 bg-white/5 p-4"
            >
              <div>
                <div className="flex items-start justify-between gap-2">
                  <h3 className="text-xl font-bold tracking-tight text-white">
                    {record.title}
                  </h3>
                  {!isLoading && (
                    <div className="flex items-center gap-1.5 text-white/70">
                      {canView ? (
                        <Eye className="size-4" aria-label="Can view" />
                      ) : (
                        <EyeOff className="size-4" aria-label="Cannot view" />
                      )}
                      {canEdit && (
                        <Pencil className="size-4" aria-label="Can edit" />
                      )}
                    </div>
                  )}
                </div>
                <Spacer size="small" />
                <p className="text-sm italic leading-relaxed text-white/60">
                  {record.description}
                </p>
              </div>
              <Spacer size="small" />
              <div className="inline-flex items-center self-start border border-white/20 bg-white/10 px-2.5 py-1">
                <span className="text-[10px] font-semibold uppercase tracking-wider text-white/90">
                  {getOwner(accounts, record.account_id)}
                </span>
              </div>
            </li>
          );
        })}
      </ul>
    </section>
  );
};

export default MedicalRecordsSection;
