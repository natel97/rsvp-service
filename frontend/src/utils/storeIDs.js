const localStorageKey = "invitations";

export const getExistingInvitations = () => {
  const stored = window.localStorage.getItem(localStorageKey);
  if (!stored) {
    return [];
  }

  const storedList = JSON.parse(stored);
  if (!storedList.length) {
    return [];
  }

  return storedList;
};

export const storeInvitation = (id) => {
  const existing = getExistingInvitations();
  for (const val of existing) {
    if (val === id) {
      return;
    }
  }

  window.localStorage.setItem(
    localStorageKey,
    JSON.stringify([...existing, id])
  );
};
