import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import MedicalRecordsSection from "#/components/(sections)/medical-records-section/medical-records-section";
import AccountSection from "#/components/(sections)/account-section/account-section";
import { GetAccounts } from "#/services/accounts/accounts.queries";
import { GetMedicalRecords } from "#/services/medical-records/medical-records.queries";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { useAccount } from "#/stores/account-store/account-store";

export const Route = createFileRoute("/(pages)/dashboard")({
  component: RouteComponent,
});

function RouteComponent() {
  const { account } = useAccount();

  const { data: accounts = [] } = useQuery(GetAccounts());
  const { data: medicalRecords = [] } = useQuery(GetMedicalRecords());

  return (
    <PageLayout>
      <AccountSection accounts={accounts} />
      {account && (
        <>
          <Spacer size="medium" />
          <MedicalRecordsSection
            accounts={accounts}
            medicalRecords={medicalRecords}
          />
        </>
      )}
    </PageLayout>
  );
}
