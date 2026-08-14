import { ArrowRightLeft } from "lucide-react";
import { useState } from "react";

import SelectAccountModal from "#/components/(modals)/select-account-modal/select-account-modal";
import { buttonVariants } from "#/components/ui/button";
import type { Account } from "#/domain/accounts";
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
          <div>
            <p className="text-xl font-semibold">
              {account.first_name} {account.last_name}
            </p>
          </div>
          <button
            type="button"
            onClick={() => setOpen(true)}
            className={cn(
              buttonVariants({ variant: "outline", size: "icon-lg" }),
              "cursor-pointer gap-2 border-2 border-dashed border-white px-4",
            )}
          >
            <ArrowRightLeft className="h-4 w-4" />
          </button>
        </div>
      ) : (
        <button
          type="button"
          onClick={() => setOpen(true)}
          className={cn(
            buttonVariants({ variant: "outline", size: "lg" }),
            "h-12 cursor-pointer border-2 border-dashed border-white px-6 text-lg",
          )}
        >
          Select Account
        </button>
      )}

      <SelectAccountModal
        accounts={accounts}
        open={open}
        onOpenChange={setOpen}
        onSelect={handleSelect}
      />
    </section>
  );
};

export default AccountSection;
