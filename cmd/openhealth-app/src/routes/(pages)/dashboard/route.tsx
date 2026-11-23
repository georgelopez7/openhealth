import { useQuery } from "@tanstack/react-query";
import { createFileRoute, Link } from "@tanstack/react-router";

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
          <div className="grid w-full grid-cols-1 items-start gap-6 border border-neutral-700 bg-black font-retro px-12 py-8 text-left sm:px-16 sm:py-12 lg:grid-cols-[20rem_minmax(0,1fr)]">
            <Link
              to="/"
              className="col-span-full inline-flex w-fit items-center gap-2 border border-neutral-700 bg-[#101010] px-3 py-2 text-xs text-white transition-colors hover:border-neutral-500 hover:bg-neutral-900"
            >
              <svg
                aria-hidden="true"
                className="size-4"
                fill="currentColor"
                shapeRendering="crispEdges"
                viewBox="0 0 16 16"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path d="M1 7h2V5h2V3h2V1h2v2h2v2h2v2h2v2h-2v6H9v-4H7v4H3V9H1V7Z" />
              </svg>
              Home
            </Link>
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
