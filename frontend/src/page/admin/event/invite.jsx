import { useNavigate } from "react-router-dom";
import { ActionButton, PageHeader } from "../../../components";

const Invite = () => {
  const navigate = useNavigate();

  return (
    <div>
      <PageHeader>Invite to Event</PageHeader>
      <ActionButton onClick={() => navigate("./person")}>
        Invite Person
      </ActionButton>
      <ActionButton onClick={() => navigate("./group")}>
        Invite Group
      </ActionButton>
    </div>
  );
};

export default Invite;
