export const DEFAULT_OG_IMAGE = "/og-image.png";

export const getSiteURL = (): string => {
  return import.meta.env.VITE_SITE_URL ?? "http://localhost:3000";
};

export const getOGImage = () => {
  return `${getSiteURL()}${DEFAULT_OG_IMAGE}`;
};
