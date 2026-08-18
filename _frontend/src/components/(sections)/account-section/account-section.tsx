import { ArrowRightLeft } from "lucide-react";
import { useState } from "react";

import SelectAccountModal from "#/components/(modals)/select-account-modal/select-account-modal";
import type { Account } from "#/domain/accounts";
import { getAccountDisplayName } from "#/domain/accounts";
import { AccountIcon } from "#/domain/icon";
import { cn } from "#/lib/utils";
import { useAccount } from "#/stores/account-store/account-store";

interface IProps {
  accounts: Account[];
}

const AccountSection = ({ accounts }: IProps) => {
  const [open, setOpen] = useState(false);
  const { account, setAccount } = useAccount();

  const handleSelect = (selectedAccount: Account) => {
    setAccount(selectedAccount);
    setOpen(false);
  };

  return (
    <section className="w-full">
      {account ? (
        <div className="flex w-full items-center justify-between py-2 text-white">
          <div className="flex h-12 items-center gap-3 rounded-lg border border-white/20 bg-white/5 px-4 py-2.5">
            <AccountIcon avatar={account.avatar} className="size-8 shrink-0" />
            <p className="text-sm font-semibold md:text-base">
              {getAccountDisplayName(account)}
            </p>
          </div>
          <button
            type="button"
            onClick={() => setOpen(true)}
            className={cn(
              "inline-flex h-12 items-center justify-center gap-2 rounded-lg border border-white/20 bg-white/5 px-4 text-sm font-medium text-white transition-colors hover:bg-white/10 md:text-base cursor-pointer",
            )}
          >
            <ArrowRightLeft className="size-6 shrink-0" />
            <span className="hidden md:inline">Switch</span>
          </button>
        </div>
      ) : (
        <button
          type="button"
          onClick={() => setOpen(true)}
          className={cn(
            "inline-flex h-12 w-full cursor-pointer items-center justify-center gap-2 rounded-lg border border-white/20 bg-white/5 px-6 text-lg font-medium text-white transition-colors hover:bg-white/10",
          )}
        >
          Select Account
        </button>
      )}
      <SelectAccountModal
        accounts={accounts}
        selectedAccountID={account?.id}
        open={open}
        onOpenChange={setOpen}
        onSelect={handleSelect}
      />
    </section>
  );
};

export default AccountSection;
