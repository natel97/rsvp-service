import { Action, EventCard, PageHeader } from "../components";

const Invitation = () => {
  return (
    <div
      className="full-width full-height"
      style={{ display: "flex", flexDirection: "column" }}
    >
      <div className="full-width full-height">
        <PageHeader>You Are Invited</PageHeader>
        <EventCard />
      </div>
      <Action>Add to Calendar</Action>
      <Action>RSVP (No Response)</Action>
      {/* <Action>Get Reminder</Action> */}
    </div>
  );
};

export default Invitation;
