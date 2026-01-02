const SummaryCard = ({title, total}: { title: string, total: string }) => {
  // const colorVariant = {
  //   warning: "text-yellow-500",
  //   error: "text-orange-500/80",
  //   debug: "text-blue-700/80",
  //   fatal: "text-red-900/80",
  // }

  return (
      <div className={"card lift flex flex-col gap-2 w-1/2 h-1/2 min-h-[200px] min-w-[150px] text-center"}>
        <span
            className={`font-body text-black  font-bold text-2xl text-shadow-white/25 text-shadow-sm`}>
          {title.toUpperCase()}
        </span>
        <div
            className={"flex justify-center items-center text-7xl text-black w-full h-full  bg-white/50 rounded-lg "}>
          {total}
        </div>

      </div>
  )
}

export default SummaryCard