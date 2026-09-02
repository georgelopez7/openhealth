import { ArrowRightLeft } from "lucide-react";
import { useState } from "react";
import SelectAccountModal from "#/components/(modals)/select-account-modal/select-account-modal";
import { Button } from "#/components/ui/button";
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
        <div className="flex w-full justify-center py-2 text-white">
          <div className="flex h-12 w-full items-center justify-between gap-2 rounded-lg border border-white/20 bg-white/5 px-2 py-2.5 sm:w-auto sm:justify-center sm:gap-10">
            <div className="flex items-center gap-2">
              <AccountIcon
                avatar={account.avatar}
                className="size-9 shrink-0"
              />
              <p className="text-sm font-semibold md:text-lg">
                {getAccountDisplayName(account)}
              </p>
            </div>
            <Button
              variant="ghost"
              size="icon"
              type="button"
              onClick={() => setOpen(true)}
              aria-label="Switch account"
              className="cursor-pointer text-white hover:text-white"
            >
              <ArrowRightLeft className="size-5" />
            </Button>
          </div>
        </div>
      ) : (
        <button
          type="button"
          onClick={() => setOpen(true)}
          className={cn(
            "inline-flex h-12 w-full cursor-pointer items-center justify-center gap-2 rounded-lg border border-white/20 bg-white/5 px-6 text-base font-medium text-white transition-colors hover:bg-white/10 md:text-lg",
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
