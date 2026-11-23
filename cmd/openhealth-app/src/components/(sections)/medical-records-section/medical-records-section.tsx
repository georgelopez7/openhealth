import { useQueries } from "@tanstack/react-query";

import MedicalRecordCard from "#/components/(cards)/medical-record-card/medical-record-card";
import type { Account } from "#/domain/accounts";
import type { MedicalRecord } from "#/domain/medical-records";
import { GetMedicalRecordAccess } from "#/services/medical-records/medical-records.queries";

interface IProps {
  account: Account;
  accounts: Account[];
  medicalRecords: MedicalRecord[];
}

// getOwner - Returns the owner's account or undefined if not found.
const getOwner = (
  accounts: Account[],
  accountID: string,
): Account | undefined => {
  return accounts.find((account) => account.id === accountID);
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
      <h2 className="text-lg font-bold md:text-xl">Medical Records</h2>
      <ul className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {medicalRecords.map((record, index) => {
          const { canView, canEdit, isLoading } = accessResults[index];
          const owner = getOwner(accounts, record.account_id);

          return (
            <MedicalRecordCard
              key={record.id}
              record={record}
              owner={owner}
              canView={canView}
              canEdit={canEdit}
              isLoading={isLoading}
            />
          );
        })}
      </ul>
    </section>
  );
};

export default MedicalRecordsSection;
