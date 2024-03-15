import { ActionButton } from "../components";
import { getAuthToken } from "../page/admin/auth";

const urlBase64ToUint8Array = (base64String) => {
  const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
  const base64 = (base64String + padding).replace(/-/g, "+").replace(/_/g, "/");

  const rawData = atob(base64);
  const outputArray = new Uint8Array(rawData.length);

  for (let i = 0; i < rawData.length; ++i) {
    outputArray[i] = rawData.charCodeAt(i);
  }

  return outputArray;
};

const saveSubscription = async (url, subscription) => {
  const apiKey = getAuthToken();
  const response = await fetch(url, {
    method: "post",
    headers: { Authorization: apiKey },
    body: JSON.stringify({
      subscription: JSON.stringify(subscription),
    }),
  });

  return response.text();
};

const subscribe = (url) =>
  navigator.serviceWorker?.getRegistration().then((reg) =>
    reg.pushManager
      .subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(
          import.meta.env.VITE_VAPID_PUBLIC_KEY
        ),
      })
      .then((val) => saveSubscription(url, val))
  );

const checkPermission = () => {
  if (!("serviceWorker" in navigator)) {
    return false;
  }

  if (!("Notification" in window)) {
    return false;
  }

  if (!("PushManager" in window)) {
    return false;
  }

  return true;
};

const requestNotificationPermission = async () => {
  const permission = await Notification.requestPermission();

  if (permission !== "granted") {
    throw new Error("Notification permission not granted");
  }
};

export const NotifyButton = ({ url }) => {
  const pushSupported = checkPermission();
  if (!pushSupported) {
    return <div>Push Notifications Not Supported</div>;
  }

  return (
    <ActionButton
      onClick={() => requestNotificationPermission().then(() => subscribe(url))}
    >
      Enable Notifications 🔔
    </ActionButton>
  );
};
