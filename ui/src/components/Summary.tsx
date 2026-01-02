import SummaryCard from "./SummaryCard.tsx";

type CardProps = {
  level: string
  total: string
}

const Summary = () => {
  // initial state
  const logSummaries: CardProps[] = [
    {level: "error", total: "0",},
    {level: "warning", total: "0",},
    {level: "debug", total: "0",},
    {level: "fatal", total: "0",}]


  // fetch array of logs from api

  // unpack into state
  // const [logLevels, setLogLevels] = useState<CardProps[]>(logSummaries);


  return (
      <div className={"card justify-center gap-2"}>
        <span className={"font-body text-3xl"}>Summary</span>
        <div className={"pt-3"}>
          <div className={"grid grid-cols-2  lg:grid-cols-4 gap-4"}>
            {logSummaries.map((log) => (
                <SummaryCard title={log.level} total={log.total}/>
            ))}
          </div>
        </div>
      </div>
  )
}

export default Summary