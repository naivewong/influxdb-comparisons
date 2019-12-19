import json
import random

if __name__ == "__main__":
	measurements = []
	with open("node_exporter.json", "r") as f:
		line = f.readline()
		while line:
			data = json.loads(line.strip())
			measurements.append(data["__name__"])
			line = f.readline()
	with open("type1.txt", "w") as f:
		chosen = set()
		for i in range(200):
			c = random.randint(0, len(measurements) - 1)
			while measurements[c] in chosen:
				c = random.randint(0, len(measurements) - 1)
			chosen.add(c)
			f.write(measurements[c] + "\n")