import { useQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
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

  const { data: accounts = [], isLoading: isLoadingAccounts } = useQuery(
    GetAccounts(),
  );
  const { data: medicalRecords = [] } = useQuery(GetMedicalRecords());

  return (
    <PageLayout navbar={false}>
      <div className="flex flex-1 items-center justify-center">
        <div className="flex w-full bg-[#101010] p-3 sm:p-4">
          <div className="grid w-full grid-cols-1 items-start gap-6 border border-neutral-700 bg-black font-retro px-12 py-14 text-left sm:px-16 sm:py-20 lg:grid-cols-[20rem_minmax(0,1fr)]">
            <AccountSection accounts={accounts} isLoading={isLoadingAccounts} />
            {account ? (
              <MedicalRecordsSection
                account={account}
                accounts={accounts}
                medicalRecords={medicalRecords}
              />
            ) : (
              <div className="flex min-h-40 items-center justify-center border border-dashed border-neutral-800 px-4 text-center text-xs text-neutral-500">
                Select a character to view medical records.
              </div>
            )}
          </div>
        </div>
      </div>
    </PageLayout>
  );
}
