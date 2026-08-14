import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { ArrowRightLeft } from "lucide-react";
import { useState } from "react";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import SelectCharacterModal from "#/components/(modals)/select-character-modal/select-character-modal";
import { buttonVariants } from "#/components/ui/button";
import type { Account } from "#/domain/accounts";
import { cn } from "#/lib/utils";
import { GetAccounts } from "#/services/accounts/accounts.queries";
import { useAccount } from "#/stores/account-store/account-store";

export const Route = createFileRoute("/(pages)/dashboard")({
  component: RouteComponent,
});

function RouteComponent() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const { data: accounts = [] } = useQuery(GetAccounts());
  const { account, setAccount } = useAccount();

  console.log("Accounts:", accounts);

  const handleSelect = (account: Account) => {
    console.log("Selected account:", account);
    setAccount(account);
    setIsModalOpen(false);
  };

  return (
    <PageLayout>
      {account ? (
        <div className="flex w-full items-center justify-between py-2 text-white">
          <div>
            <p className="text-lg font-semibold">
              {account.first_name} {account.last_name}
            </p>
            <p className="text-sm opacity-80">{account.email}</p>
          </div>
          <button
            type="button"
            onClick={() => setIsModalOpen(true)}
            className={cn(
              buttonVariants({ variant: "outline", size: "sm" }),
              "cursor-pointer gap-2 border-2 border-dashed border-white px-4",
            )}
          >
            <ArrowRightLeft className="h-4 w-4" />
            Switch
          </button>
        </div>
      ) : (
        <button
          type="button"
          onClick={() => setIsModalOpen(true)}
          className={cn(
            buttonVariants({ variant: "outline", size: "lg" }),
            "h-12 cursor-pointer border-2 border-dashed border-white px-6 text-lg",
          )}
        >
          Select Character
        </button>
      )}
      <SelectCharacterModal
        accounts={accounts}
        open={isModalOpen}
        onOpenChange={setIsModalOpen}
        onSelect={handleSelect}
      />
    </PageLayout>
  );
}
