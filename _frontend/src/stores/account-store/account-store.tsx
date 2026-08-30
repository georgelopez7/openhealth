import { create } from "zustand";
import type { Account } from "#/domain/accounts";

interface IAccountStore {
  account: Account | null;
  setAccount: (account: Account) => void;
}

export const useAccount = create<IAccountStore>((set) => ({
  account: null,
  setAccount: (account) => set({ account }),
}));
