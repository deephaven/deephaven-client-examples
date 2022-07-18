import pydeephaven
from pydeephaven import Session
from pydeephaven import ComboAggregation
from pydeephaven import DHError

try:
	# The sync_fields argument has no equivalent in C++,
	# and is instead two different methods in Go.
	# Also note that this is called "Session" instead of "Client".
	# Also, should a port be a string or a number? Should the host and port be two different arguments?
	client = Session("localhost", 10000, sync_fields=pydeephaven.session.SYNC_ONCE)

	# Creation of a pyarrow or pandas record omitted,
	# since that's always going to be different from other languages...
	some_record = None

	handle1 = client.import_table(some_record)

	handle2 = client.empty_table(10)

	handle3 = handle1.where(["a > 10", "b % 2 == 0"])

	handle4 = client.merge_tables([handle1, handle3], order_by="a")

	combo_agg = (ComboAggregation()
		.sum(cols=["foo = a", "bar = c"])
		.avg(cols=["avgA = a", "avgB = b"]))
	handle5 = handle4.agg_by(combo_agg, ["g"])

	snapshot_record = handle5.snapshot()
	print(snapshot_record)
except DHError as e:
	print("Caught an exception:", e)