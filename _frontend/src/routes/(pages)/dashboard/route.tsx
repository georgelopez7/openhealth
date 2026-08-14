import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { useState } from "react";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import SelectCharacterModal from "#/components/(modals)/select-character-modal/select-character-modal";
import { buttonVariants } from "#/components/ui/button";
import type { Account } from "#/domain/accounts";
import { cn } from "#/lib/utils";
import { GetAccounts } from "#/services/accounts/accounts.queries";

export const Route = createFileRoute("/(pages)/dashboard")({
  component: RouteComponent,
});

function RouteComponent() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const { data: accounts = [] } = useQuery(GetAccounts());

  console.log("Accounts:", accounts);

  const handleSelect = (account: Account) => {
    console.log("Selected account:", account);
    setIsModalOpen(false);
  };

  return (
    <PageLayout>
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
      <SelectCharacterModal
        accounts={accounts}
        open={isModalOpen}
        onOpenChange={setIsModalOpen}
        onSelect={handleSelect}
      />
    </PageLayout>
  );
}
