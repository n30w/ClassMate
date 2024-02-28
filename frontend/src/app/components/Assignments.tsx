const Assignments = () => {
    const homeworks = [
        {'name': 'Homework 1', 'duedate': 'Feb 8'},
        {'name': 'Homework 2', 'duedate': 'Feb 25'},
        {'name': 'Homework 3', 'duedate': 'March 3'},
        {'name': 'Homework 4', 'duedate': 'March 10'}
    ]

    const hwElements = homeworks.map(homework => {
        return (
            <li>
                <h2 class="text-m font-bold">{homework.name}</h2>
                <p class="text-sm font-light pb-1 mb-1 border-b border-black">{homework.duedate}</p>
            </li>
        )
    })

  return (
    <div class="w-60 h-70 px-4">
        <h1 class="font-bold text-xl border-b-2 border-black mb-4 pb-4">Assignments</h1>
        <ul>{hwElements}</ul>
    </div>
  )
}

export default Assignments