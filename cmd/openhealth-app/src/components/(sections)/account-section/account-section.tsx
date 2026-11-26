import { Check, LoaderCircle } from "lucide-react";
import type { Account } from "#/domain/accounts";
import { getAccountDisplayName } from "#/domain/accounts";
import { AccountIcon } from "#/domain/icon";
import { cn } from "#/lib/utils";
import { useAccount } from "#/stores/account-store/account-store";

interface IProps {
  accounts: Account[];
  isLoading: boolean;
}

const AccountSection = ({ accounts, isLoading }: IProps) => {
  const { account, setAccount } = useAccount();

  return (
    <section className="w-full lg:w-80 lg:shrink-0">
      <div className="border border-neutral-700 bg-[#101010] p-4 sm:p-5">
        <div className="mb-4 border-b border-neutral-800 pb-3">
          <p className="truncate text-xs text-white">
            {account
              ? getAccountDisplayName(account)
              : isLoading
                ? "Loading characters..."
                : "Choose an avatar"}
          </p>
        </div>

        {isLoading ? (
          <div className="flex min-h-24 items-center justify-center text-neutral-500">
            <LoaderCircle className="size-6 animate-spin" />
          </div>
        ) : accounts.length === 0 ? (
          <p className="py-4 text-center text-xs text-neutral-500">
            No characters available.
          </p>
        ) : (
          <div className="grid grid-cols-2 gap-2">
            {accounts.map((character) => {
              const isSelected = character.id === account?.id;

              return (
                <button
                  key={character.id}
                  type="button"
                  aria-pressed={isSelected}
                  onClick={() => setAccount(character)}
                  className={cn(
                    "relative flex min-h-24 min-w-0 flex-col items-center justify-center gap-2 border p-2 text-center transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-white/50",
                    isSelected
                      ? "border-neutral-400 bg-[#242424] text-white"
                      : "border-neutral-800 bg-black text-neutral-400 hover:border-neutral-600 hover:text-white",
                  )}
                >
                  {isSelected && (
                    <Check
                      aria-hidden="true"
                      className="absolute top-1.5 right-1.5 size-3 text-white"
                    />
                  )}
                  <AccountIcon
                    avatar={character.avatar}
                    className="size-10 shrink-0"
                  />
                  <span className="flex min-w-0 max-w-full flex-col break-words text-[9px] leading-relaxed">
                    <span>{character.first_name}</span>
                    <span>{character.last_name}</span>
                  </span>
                </button>
              );
            })}
          </div>
        )}
      </div>
    </section>
  );
};

export default AccountSection;
