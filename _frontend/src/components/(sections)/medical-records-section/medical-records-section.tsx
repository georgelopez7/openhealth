import Spacer from "#/components/(layouts)/spacer/spacer";
import type { Account } from "#/domain/accounts";
import type { MedicalRecord } from "#/domain/medical-records";

interface IProps {
  accounts: Account[];
  medicalRecords: MedicalRecord[];
}

// getOwner - Returns the owner's name or the account ID if the owner is not found.
const getOwner = (accounts: Account[], accountId: string): string => {
  const owner = accounts.find((account) => account.id === accountId);
  return owner ? `${owner.first_name} ${owner.last_name}` : accountId;
};

const MedicalRecordsSection = ({ accounts, medicalRecords }: IProps) => {
  if (medicalRecords.length === 0) {
    return <p className="opacity-80">No medical records found.</p>;
  }

  return (
    <section className="w-full space-y-4 text-white">
      <h2 className="text-xl font-bold">Medical Records</h2>
      <ul className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {medicalRecords.map((record) => (
          <li
            key={record.id}
            className="flex min-h-45 flex-col justify-between rounded-lg border border-white/20 bg-white/5 p-4"
          >
            <div>
              <h3 className="text-xl font-bold tracking-tight text-white">
                {record.title}
              </h3>
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
        ))}
      </ul>
    </section>
  );
};

export default MedicalRecordsSection;
