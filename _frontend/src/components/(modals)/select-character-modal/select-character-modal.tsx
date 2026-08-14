import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import type { Account } from "@/domain/accounts";
import { cn } from "@/lib/utils";

interface IProps {
  accounts: Account[];
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
  onSelect?: (account: Account) => void;
}

const SelectCharacterModal = ({
  accounts,
  open,
  onOpenChange,
  onSelect,
}: IProps) => {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Select Character</DialogTitle>
          <DialogDescription>Choose an account to continue.</DialogDescription>
        </DialogHeader>
        <div className="grid grid-cols-2 gap-4">
          {accounts.map((account) => (
            <button
              key={account.id}
              type="button"
              onClick={() => onSelect?.(account)}
              className={cn(
                "flex items-center justify-center rounded-lg border border-dashed border-white p-4 text-center cursor-pointer hover:bg-white/10",
              )}
            >
              {account.first_name} {account.last_name}
            </button>
          ))}
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default SelectCharacterModal;
