import Spacer from "#/components/(layouts)/spacer/spacer";
import type { Account } from "#/domain/accounts";
import { getAccountDisplayName } from "#/domain/accounts";
import { AccountIcon } from "#/domain/icon";
import type { MedicalRecord } from "#/domain/medical-records";
import { cn } from "#/lib/utils";

interface IPermissionBannerProps {
  action: "view" | "edit";
  granted: boolean;
}

const PermissionBanner = ({ action, granted }: IPermissionBannerProps) => {
  return (
    <span
      className={cn(
        "border px-2 py-0.5 text-[10px] font-semibold uppercase tracking-wider",
        granted
          ? "border-white/20 bg-white/10 text-white"
          : "border-white/10 bg-white/3 text-white/30 opacity-40 cursor-not-allowed",
      )}
    >
      Can {action}
    </span>
  );
};

interface IProps {
  record: MedicalRecord;
  owner?: Account;
  canView: boolean;
  canEdit: boolean;
  isLoading: boolean;
}

const MedicalRecordCard = ({
  record,
  owner,
  canView,
  canEdit,
  isLoading,
}: IProps) => {
  const ownerName = owner ? getAccountDisplayName(owner) : record.account_id;

  return (
    <li className="flex min-h-45 flex-col justify-between rounded-lg border border-white/20 bg-white/5 p-4">
      <div>
        <div className="flex items-start justify-between gap-2">
          <h3 className="text-base font-bold tracking-tight text-white md:text-xl">
            {record.title}
          </h3>
          {owner && (
            <AccountIcon avatar={owner.avatar} className="size-12 shrink-0" />
          )}
        </div>
        <Spacer size="small" />
        <p className="text-[12px] italic leading-relaxed text-white/60">
          {record.description}
        </p>
      </div>
      <Spacer size="small" />
      <div className="flex flex-col items-center gap-2">
        <div className="inline-flex items-center border border-white/20 bg-white/10 px-2.5 py-1">
          <span className="text-[10px] font-semibold uppercase tracking-wider text-white/90">
            {ownerName}
          </span>
        </div>
        {!isLoading && (
          <div className="flex items-center gap-2">
            <PermissionBanner action="view" granted={canView} />
            <PermissionBanner action="edit" granted={canEdit} />
          </div>
        )}
      </div>
    </li>
  );
};

export default MedicalRecordCard;
