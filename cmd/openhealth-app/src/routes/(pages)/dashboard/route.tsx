import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import AccountSection from "#/components/(sections)/account-section/account-section";
import MedicalRecordsSection from "#/components/(sections)/medical-records-section/medical-records-section";
import { getOGImage, getSiteURL } from "#/lib/seo";
import { GetAccounts } from "#/services/accounts/accounts.queries";
import { GetMedicalRecords } from "#/services/medical-records/medical-records.queries";
import { useAccount } from "#/stores/account-store/account-store";

const SEO = {
  url: getSiteURL(),
  title: "Dashboard | OpenHealth",
  description:
    "View and manage your health accounts and medical records powered by OpenFGA authorization.",
  image: getOGImage(),
};

export const Route = createFileRoute("/(pages)/dashboard")({
  head: () => ({
    meta: [
      {
        title: SEO.title,
      },
      {
        name: "description",
        content: SEO.description,
      },
      {
        property: "og:title",
        content: SEO.title,
      },
      {
        property: "og:description",
        content: SEO.description,
      },
      {
        property: "og:type",
        content: "website",
      },
      {
        property: "og:image",
        content: SEO.image,
      },
      {
        name: "twitter:card",
        content: "summary_large_image",
      },
      {
        name: "twitter:title",
        content: SEO.title,
      },
      {
        name: "twitter:description",
        content: SEO.description,
      },
      {
        name: "twitter:image",
        content: SEO.image,
      },
    ],
  }),
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
            account={account}
            accounts={accounts}
            medicalRecords={medicalRecords}
          />
        </>
      )}
    </PageLayout>
  );
}
