import { styled } from "styled-components";
import moment from "moment";
import { useState } from "react";

const colors = {
  success: {
    light: "#AF7",
    dark: "#053",
  },
  primary: {
    light: "#6CC",
    dark: "#088",
  },
  secondary: {
    light: "#BDF",
    dark: "#046",
  },
  background: {
    light: "#DDD",
    dark: "#222",
  },
};

const sizes = {
  sm: "18px",
  md: "22px",
  lg: "28px",
};

const TextInputContainer = styled.div`
  margin: 12px 8px;
`;

export const TextInput = (params) => {
  return (
    <TextInputContainer>
      {params.area ? <RawTextArea {...params} /> : <RawTextInput {...params} />}
    </TextInputContainer>
  );
};

export const Modal = ({ open = false, setIsOpen = () => null, children }) => {
  if (!open) return null;

  return (
    <ModalBack onClick={() => setIsOpen(false)}>
      <ModalInner onClick={(e) => e.stopPropagation()}>{children}</ModalInner>
    </ModalBack>
  );
};

export const Text = styled.span`
  padding: 0;
  margin: 0;
  color: white;
  @media (prefers-color-scheme: light) {
    color: black;
  }
  font-size: ${(params) => sizes[params.size || "sm"]};
`;

const ModalBack = styled.div`
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: #000a;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const ModalInner = styled.div`
  background: ${colors.background.dark};
  @media (prefers-color-scheme: light) {
    background: ${colors.background.light};
  }
  margin: 16px;
  max-height: 75vh;
  overflow: auto;
`;

const commonTextStyles = `
  border: 1px solid #ccc;
  width: 100%;
  padding: 12px;
  font-size: 1.5rem;
  border-radius: 8px;`;

const RawTextInput = styled.input`
  ${commonTextStyles}
`;

const RawTextArea = styled.textarea`
  ${commonTextStyles}
`;

export const Radio = styled.input`
  & + label {
    box-shadow: 0px 0px 2px 2px rgba(255, 255, 255, 0.25);
    border-radius: 8px;
    padding: 8px;
  }

  &:checked + label {
    background: ${colors.success.dark};
  }

  &:checked + label {
    @media (prefers-color-scheme: light) {
      background: ${colors.success.light};
    }
  }
`;

export const PageHeader = styled.h1`
  background: ${colors.primary.dark};
  padding: 1.3rem;
  text-align: center;
  font-size: 38px;

  @media (prefers-color-scheme: light) {
    background: ${colors.primary.light};
  }
`;

export const ModalBackground = styled.div`
  position: absolute;
  left: 0;
  top: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.75);
  display: flex;
  align-items: center;
  justify-content: center;
`;

export const ModalBody = styled.div`
  background: #444;
  padding: 32px;
  border-radius: 8px;
  overflow: auto;
  display: flex;
  justify-content: space-between;
  flex-direction: column;

  @media (prefers-color-scheme: light) {
    background: ${colors.secondary.light};
  }
`;

export const ActionButton = styled.button`
  font-size: 1.5rem;
  background: ${colors.success.dark};
  margin: 8px 0;
  padding: 1rem;
  border-radius: 8px;

  @media (prefers-color-scheme: light) {
    background: ${colors.success.light};
  }
`;

export const Card = styled.div`
  padding: 16px;
  margin: 32px 16px;
  text-align: left;
  border-radius: 8px;
  box-shadow: 0px 0px 4px 4px rgba(0, 0, 0, 0.25);
  background: ${colors.secondary.dark};

  @media (prefers-color-scheme: light) {
    background: ${colors.secondary.light};
  }
`;

export const Option = ({
  prefix,
  options = [],
  current = "",
  fullWidth = false,
  setCurrent = () => {},
}) => {
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        margin: "16px",
        ...(fullWidth && { width: "100%" }),
      }}
    >
      {options.map((option) => (
        <span key={option}>
          <Radio
            id={`${prefix}-${option}`}
            name={prefix}
            type="radio"
            value={option}
            checked={option === current}
            onChange={() => setCurrent(option)}
            style={{ display: "none" }}
          />
          <label htmlFor={`${prefix}-${option}`}>{option}</label>
        </span>
      ))}
    </div>
  );
};

const Attendance = ({ yes, no, maybe }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        marginTop: "48px",
      }}
    >
      <div>🟢 Yes ({yes})</div>
      <div>🤷‍♂️ Maybe ({maybe})</div>
      <div>🔴 No ({no})</div>
    </div>
  );
};

const EmojiDetail = ({ emoji = "", children, size = "1rem" }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
        alignItems: "center",
        margin: "16px 0",
      }}
    >
      <div style={{ fontSize: size, marginRight: "8px" }}>{emoji}</div>
      <div>{children}</div>
    </div>
  );
};

export const DAY_FORMAT = "MMMM Do YYYY";
export const TIME_FORMAT = "hh:mm A";

export const EventCard = ({
  title = "- - -",
  date = Date.parse("23 Jan 2023 18:30:00 GMT+1100"),
  street = "- - -",
  city = "- - - -, - - -  0000",
  attendance = { yes: -1, no: -1, maybe: -1 },
  description = "-----",
  onClick,
}) => {
  const day = moment(date).format(DAY_FORMAT);
  const time = moment(date).format(TIME_FORMAT);
  const address = encodeURIComponent(street + " " + city);

  return (
    <Card onClick={onClick}>
      <h2 style={{ fontSize: "2rem", textAlign: "center" }}>{title}</h2>

      {date && (
        <EmojiDetail size="2rem" emoji="📅">
          <h3>{day}</h3>
          <h3>{time}</h3>
        </EmojiDetail>
      )}

      <EmojiDetail size="2rem" emoji="📍">
        <a
          href={`https://maps.google.com/?q=${address}`}
          target="_blank"
          rel="noreferrer"
          style={{
            color: "inherit",
            textDecoration: "none",
            alignSelf: "flex-start",
          }}
        >
          <h3>{street}</h3>
          <h3>{city}</h3>
        </a>
      </EmojiDetail>

      <div>{description}</div>
      {attendance && date && <Attendance {...attendance} />}
    </Card>
  );
};

const OptionStyle = styled.div`
  border-radius: 16px;
  padding: 8px;
  margin: 8px;
`;

const SelectionContainer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  width: 100%;
`;

export const TimeOption = ({ option, onSelect }) => {
  const day =
    moment(option.time).format(DAY_FORMAT) +
    " " +
    moment(option.time).format(TIME_FORMAT);
  const [upvote, setUpvote] = useState(option.isUpvote);
  const [downvote, setDownvote] = useState(option.isDownvote);

  const currentSet = upvote ? "Likely" : downvote ? "Unlikely" : "idk";
  const setCurrent = (val) => {
    const map = {
      Likely: true,
      idk: null,
      Unlikely: false,
    };

    onSelect(map[val]);
    setUpvote(val === "Likely");
    setDownvote(val === "Unlikely");
  };

  return (
    <Card>
      <OptionStyle>
        <Text size="md">{day}</Text>
        <SelectionContainer>
          <SelectionContainer>
            <Option
              fullWidth
              prefix={option.id}
              options={["Likely", "idk", "Unlikely"]}
              current={currentSet}
              setCurrent={setCurrent}
            />
          </SelectionContainer>
        </SelectionContainer>
      </OptionStyle>
    </Card>
  );
};
