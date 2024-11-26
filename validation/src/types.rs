bitflags::bitflags!{
	#[derive(Clone,Copy,Debug,Default)]
	pub struct Games:u32{
		const Bhop=1<<0;
		const Surf=2<<0;
		const FlyTrials=5<<0;
	}
}
