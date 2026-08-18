import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import type { Account } from "@/domain/accounts";
import { getAccountDisplayName } from "@/domain/accounts";
import { AccountIcon } from "@/domain/icon";
import { cn } from "@/lib/utils";

interface IProps {
  accounts: Account[];
  selectedAccountID?: string;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
  onSelect?: (account: Account) => void;
}

const SelectAccountModal = ({
  accounts,
  selectedAccountID,
  open,
  onOpenChange,
  onSelect,
}: IProps) => {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Select Account</DialogTitle>
          <DialogDescription className="text-xs">
            Choose an account to continue.
          </DialogDescription>
        </DialogHeader>
        <div className="grid grid-cols-2 gap-4">
          {accounts.map((account) => {
            const isSelected = account.id === selectedAccountID;
            return (
              <button
                key={account.id}
                type="button"
                disabled={isSelected}
                onClick={() => onSelect?.(account)}
                className={cn(
                  "flex cursor-pointer flex-col items-center justify-center rounded-lg border border-white/20 bg-white/5 p-4 text-center transition-colors hover:bg-white/10",
                  isSelected && "cursor-not-allowed opacity-50",
                )}
              >
                <AccountIcon
                  avatar={account.avatar}
                  className="mb-2 h-12 w-12"
                />
                {getAccountDisplayName(account)}
              </button>
            );
          })}
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default SelectAccountModal;
