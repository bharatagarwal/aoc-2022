input = File.readlines("input.txt")
total = 0
input[0..4].each do |moves|
	opponents, yours = moves.split(' ')
	if yours == 'X'
		total += 1
		puts "yours: #{yours}"
		puts "opponents: #{opponents}"
	
		if opponents == 'A'
			total += 3
		elsif opponents == 'C'
			total += 6
		end
	elsif yours == 'Y'
		puts "yours: #{yours}"
		puts "opponents: #{opponents}"
	
		total += 2

		if opponents == 'A'
			total += 6
		elsif opponents	== 'B'
			total += 3
		end
	elsif yours == 'Z'
		puts "yours: #{yours}"
		puts "opponents: #{opponents}"
	
		total += 3

		if opponents == 'B'
			total += 6
		elsif opponents == 'C'
			total += 3
		end
	end	
end


puts total