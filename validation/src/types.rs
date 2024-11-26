bitflags::bitflags!{
	#[derive(Clone,Copy,Debug,Default)]
	pub struct Games:u32{
		const Bhop=1<<0;
		const Surf=1<<1;
		const FlyTrials=1<<4;
	}
}
