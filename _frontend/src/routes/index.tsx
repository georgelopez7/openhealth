import { createFileRoute, Link } from "@tanstack/react-router";
import GitHubIcon from "#/components/(icons)/github-icon";
import OpenhealthIcon from "#/components/(icons)/openhealth-icon";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { EXTERNAL_LINKS } from "#/domain/config";
import { cn } from "#/lib/utils";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
  return (
    <PageLayout navbar={false}>
      <div className="flex flex-1 flex-col items-center justify-center">
        <div className="flex items-center gap-6">
          <OpenhealthIcon className="size-30" />
          <h1 className="text-6xl font-bold">OpenHealth</h1>
        </div>
        <Spacer size="xsmall" />
        <div className="flex gap-3">
          <Link
            to="/dashboard"
            className={cn(
              "inline-flex h-12 items-center justify-center gap-3 rounded-lg border border-white/20 bg-white/5 px-5 text-lg font-medium text-white transition-colors hover:bg-white/10",
            )}
          >
            Get Started
          </Link>
          <a
            href={EXTERNAL_LINKS["openhealth.github"]}
            target="_blank"
            rel="noopener noreferrer"
            className={cn(
              "inline-flex size-12 items-center justify-center rounded-lg border border-white/20 bg-white/5 text-white transition-colors hover:bg-white/10",
            )}
            aria-label="GitHub"
          >
            <GitHubIcon className="size-6" />
          </a>
        </div>
        <Spacer size="medium" />
        <div className="inline-flex items-center gap-2 px-4 py-2 text-[10px] text-white/60">
          <span>
            Powered by{" "}
            <a
              href={EXTERNAL_LINKS["openfga.website"]}
              target="_blank"
              rel="noopener noreferrer"
              className="font-semibold underline-offset-2 transition-colors hover:text-white hover:underline"
            >
              OpenFGA
            </a>
          </span>
        </div>
      </div>
    </PageLayout>
  );
}
