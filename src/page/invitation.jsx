import { EventCard, PageHeader } from "../components";

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
      <div>Add to Calendar</div>
      <div>RSVP (No Response)</div>
      <div>Get Reminder</div>
    </div>
  );
};

export default Invitation;
