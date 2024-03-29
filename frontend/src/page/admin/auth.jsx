import { useEffect, useState } from "react";
import { ActionButton, PageHeader, TextInput } from "../../components";
import { useNavigate } from "react-router-dom";
import { NotifyButton } from "../../components/notify";

const authTokenKey = "auth-token";

export const getAuthToken = () => {
  const val = window.localStorage.getItem(authTokenKey);
  if (!val) {
    return "";
  }

  return val;
};

const setAuthToken = (token = "") => {
  window.localStorage.setItem(authTokenKey, token);
};

const AdminAuth = () => {
  const [apiKey, setAPIKey] = useState(getAuthToken());
  const navigate = useNavigate();

  useEffect(() => {
    setAuthToken(apiKey);
  }, [apiKey]);
  return (
    <div className="full-height full-width">
      <PageHeader>Admin</PageHeader>
      <h2>API KEY: </h2>
      <TextInput onChange={(e) => setAPIKey(e.target.value)} value={apiKey} />
      <ActionButton onClick={() => navigate("event")}>
        Manage Events
      </ActionButton>
      <ActionButton onClick={() => navigate("people")}>
        Manage People
      </ActionButton>
      <ActionButton onClick={() => navigate("group")}>
        Manage Groups
      </ActionButton>
      <NotifyButton
        url={`${import.meta.env.VITE_API_URL}/admin/notify/subscribe`}
      />
    </div>
  );
};

export default AdminAuth;
